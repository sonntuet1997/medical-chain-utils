package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
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

func EncryptMessage(message []byte, secKey []byte) ([]byte, error) {
	block, _ := aes.NewCipher(Hash256(secKey))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, message, nil)
	return ciphertext, nil
}
func DecryptCipher(data []byte, secKey []byte) ([]byte, error) {
	key := Hash256(secKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func ConvertMessage(message interface{}) ([]byte, error) {
	var bmsg []byte
	var err error
	switch message := message.(type) {
	case json.RawMessage:
		bmsg = message

	case string:
		bmsg = []byte(message)
	default:
		bmsg, err = json.Marshal(message)
		if err != nil {
			return nil, err
		}
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
func GenAuthorization(id string, publicKey string) (string, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}
	type certificateStruct struct {
		ID        string `json:"id"`
		Timestamp int64  `json:"timestamp"`
		Exp       int    `json:"exp"`
	}
	type authForm struct {
		Signature         string `json:"signature"`
		certificateStruct `json:"certificateInfo"`
		PublicKey         string `json:"publicKey"`
	}
	var auth authForm
	auth.PublicKey = publicKey
	var cert certificateStruct
	cert.ID = ID.String()
	cert.Timestamp = time.Now().Unix()
	cert.Exp = 24 * 60 * 60 * 60
	auth.Signature = publicKey
	auth.certificateStruct = cert
	str, err := json.Marshal(auth)

	if err != nil {
		return "", err
	}
	return string(str), nil
}
