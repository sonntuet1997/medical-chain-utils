package cryptography

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/google/uuid"
)

func copyBytes(b [32]byte) []byte {
	h := make([]byte, 0)
	for _, e := range b {
		h = append(h, e)
	}
	return h
}
func Hash256(a []byte) []byte {
	hash := copyBytes(sha256.Sum256(a))
	return hash
}

func ConvertMessage(message interface{}) ([]byte, error) {
	bmsg, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	hash := Hash256(bmsg)
	return hash, nil
}
func ConvertBase64ToBytes(b64 string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ConvertBytesToBase64(b []byte) string {
	b64 := base64.StdEncoding.EncodeToString(b)
	return b64
}

func GenerateKeyPair(privateKey []byte) (privkey []byte, pubkey []byte, err error) {
	if privateKey == nil {
		key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		pubkey = elliptic.MarshalCompressed(secp256k1.S256(), key.X, key.Y)
		privkey = make([]byte, 32)
		blob := key.D.Bytes()
		copy(privkey[32-len(blob):], blob)

		return privkey, pubkey, nil
	}
	var e ecdsa.PrivateKey
	e.D = new(big.Int).SetBytes(privateKey)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return e.D.Bytes(), elliptic.MarshalCompressed(secp256k1.S256(), e.X, e.Y), nil
}

func SignMessage(msg interface{}, secKey []byte) ([]byte, error) {
	hash, err := ConvertMessage(msg)
	if err != nil {
		return nil, err
	}
	sig, err := secp256k1.Sign(hash, secKey)
	if err != nil {
		log.Fatal((err))
	}
	return sig[0:64], nil
}

func VerifySig(msg interface{}, sig []byte, pub []byte) (bool, error) {
	hash, err := ConvertMessage(msg)
	if err != nil {
		return false, err
	}
	ok := secp256k1.VerifySignature(pub, hash, sig)
	if !ok {
		return false, errors.New("verify signature fail")
	}

	return ok, nil

}
func GenAuthorization(id string, privateKey string) (string, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}
	type certificateStruct struct {
		ID        string `json:"id"`
		Timestamp string `json:"timestamp"`
		Exp       int    `json:"exp"`
	}
	type authForm struct {
		Signature         string `json:"signature"`
		certificateStruct `json:"certificateInfo"`
		PublicKey         string `json:"publicKey"`
	}
	// Gen key pair
	var pk, Pk []byte
	if privateKey != "" {
		pk, Pk, err = GenerateKeyPair([]byte(privateKey))
		if err != nil {
			return "", err
		}
	} else {
		bpk, _ := ConvertBase64ToBytes("GmQE4ZljJ5PCXev2dRPCW2JHVefgsTM6+96CmqJjb0w=")
		pk, Pk, err = GenerateKeyPair(bpk)
		if err != nil {
			return "", err
		}
	}
	var auth authForm
	auth.PublicKey = ConvertBytesToBase64(Pk)
	var cert certificateStruct
	cert.ID = ID.String()
	cert.Timestamp = time.Now().String()
	cert.Exp = 24 * 60 * 60 * 60
	bsig, _ := SignMessage(cert, pk)
	sig := ConvertBytesToBase64(bsig)
	auth.Signature = sig
	auth.certificateStruct = cert
	str, err := json.Marshal(auth)

	if err != nil {
		return "", err
	}
	return string(str), nil
}
