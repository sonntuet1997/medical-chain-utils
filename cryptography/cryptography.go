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

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

func GenerateKeyPair() (privkey []byte, pubkey []byte, err error) {
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
