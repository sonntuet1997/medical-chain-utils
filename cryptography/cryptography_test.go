package cryptography

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

type TestStruct struct {
	Message    json.RawMessage `json:"message"`
	PrivateKey string          `json:"privateKey"`
	Signature  string          `json:"signature"`
}

var TESTCASE = []interface{}{
	`{
		"privateKey": "GmQE4ZljJ5PCXev2dRPCW2JHVefgsTM6+96CmqJjb0w=",
		"signature": "vGoL6rYnmA2lh4ZffBfovqOdGkiHEJ/YFvNlLfa058o9LoUx0blzjQzCmmoQ8XsOAa64bQpNMrGzsngkbYumBA==",
		"message": {
			"id": "5ef65976-d385-4525-add1-1cd2bcc6fd59",
			"timestamp": "2021-09-29T04:53:23.825Z",
			"exp": 2799360000000
		}
	}`,
	`{
		"privateKey": "GmQE4ZljJ5PCXev2dRPCW2JHVefgsTM6+96CmqJjb0w=",
		"signature": "mZW2Vtm9Pf/1QRIQpzqhgaJp9OTiV/dbKDTZC36Ig9cpqGnO86NuzmYhsigUpFYb4Bo/9DO2DDWYdSskdOTiBg==",
		"message": {
		"id": "91413875-ac0a-4aa9-a370-facfcb06cd27",
		"timestamp": "2021-09-29T04:54:23.625Z",
		"exp": 2799360000000
		}
	}`,
}

func TestSignMessage2(t *testing.T) {
	for _, k := range TESTCASE {
		var c TestStruct
		err := json.Unmarshal([]byte(k.(string)), &c)
		if err != nil {
			t.Error(err)
		}
		log.Println(c.PrivateKey)
		pk, _ := ConvertBase64ToBytes(c.PrivateKey)
		// prk := secp256k1.PrivKey{Key: pk}
		sig, err := SignMessage2(c.Message, pk)
		log.Println(len(sig))
		if err != nil {
			t.Error(err)
		}
		sig_, _ := ConvertBase64ToBytes(c.Signature)
		if !bytes.Equal(sig, sig_) {
			t.Error("not match")
		}
	}
}

func TestVerifyMessage2(t *testing.T) {
	for _, k := range TESTCASE {
		var c TestStruct
		json.Unmarshal([]byte(k.(string)), &c)

		pk, err := ConvertBase64ToBytes(c.PrivateKey)
		if err != nil {
			t.Error(err)
		}
		sig_, err := ConvertBase64ToBytes(c.Signature)
		if err != nil {
			t.Error(err)
		}
		log.Println(pk)
		if ok, _ := VerifySig2(c.Message, sig_, getPublicKey(pk)); !ok {
			t.Error("not match")
		}
	}
}

func TestRaw(t *testing.T) {
	for _, k := range TESTCASE {
		var c TestStruct
		json.Unmarshal([]byte(k.(string)), &c)
		pk, err := ConvertBase64ToBytes(c.PrivateKey)
		if err != nil {
			t.Error(err)
		}
		log.Println(pk)

		privKey := secp256k1.PrivKey{Key: pk}
		log.Println(privKey.Bytes())
		mess, _ := c.Message.MarshalJSON()
		sig, err := RawSignMessage(mess, privKey.Bytes())
		if err != nil {
			t.Error(err)
		}
		if RawVerifyMessage(mess, sig, privKey.PubKey().Bytes()) {
			t.Error("not match")
		}
	}
}

func TestGenKey(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()

	log.Println(ConvertBytesToBase64(priv.Bytes()))
	log.Println(ConvertBytesToBase64(pub.Bytes()))
	t.Error("123")
}
