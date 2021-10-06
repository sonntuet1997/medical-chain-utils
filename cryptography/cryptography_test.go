package cryptography

import (
	"encoding/json"
	"testing"
)

type TestStruct struct {
	Message    json.RawMessage `json:"message"`
	PrivateKey string          `json:"privateKey"`
	Signature  string          `json:"signature"`
}

var secKey = "12313513123123131231131cdfavdhfjanmfknflvjkzflzcjfzlknvjkzd"
var data = "123456"

func TestEncrypt(t *testing.T) {
	out, err := EncryptMessage([]byte(data), []byte(secKey))
	if err != nil {
		t.Error(err)
	}
	t.Error("123")
	t.Log(ConvertBytesToBase64(out))
	if plain, err := DecryptCipher(out, []byte(secKey)); err == nil {
		if string(plain) != data {
			t.Error("not match")
		}
		t.Error(string(plain))
	}
}
