package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sonntuet1997/medical-chain-utils/cryptography"
)

var JS = `{"message":"{\"Creator\": \"\",\"ServiceId\": \"b6f64900-2585-11ec-b8ef-128921e0ca39\",\"UserId\": \"51177332-25ae-11ec-9088-7aa190eaaaed\",\"ServiceUserId\": \"\",\"IsActive\": true}","signature":"bAuUk2R8QTt+W2G4LOucBhQPoGLdhbKQmxdhHVMfnwkTS0E71ks0HZU3G2Z1grZmTPpOvoP2vSIvEhbVTfcASA==","pubKey":"Ag316zc4okhuTy9wwb2/mnUvx5opTd7EdHPOeo+yIUIq"}`

var test = `{"message":"eyJjcmVhdG9yIjoiIiwic2VydmljZUlkIjoiYjZmNjQ5MDAtMjU4NS0xMWVjLWI4ZWYtMTI4OTIxZTBjYTM5IiwidXNlcklkIjoiNTExNzczMzItMjVhZS0xMWVjLTkwODgtN2FhMTkwZWFhYWVkIiwic2VydmljZVVzZXJJZCI6IiIsImlzQWN0aXZlIjp0cnVlfQ","signature":"tCNdfD52P1J1V0XguXm4lADJL8U8OzGNFzBCIZ0v1P4ORmtqVGHBpg3kIHhq+0V6bcIxuJ5MDN06stPvCE954Q==","pubKey":"Ag316zc4okhuTy9wwb2/mnUvx5opTd7EdHPOeo+yIUIq"}`

func main() {
	type SubMessage struct {
		Creator       string `json:"Creator"`
		ServiceId     string `json:"ServiceId"`
		UserId        string `json:"UserId"`
		ServiceUserId string `json:"ServiceUserId"`
		IsActive      bool   `json:"IsActive"`
	}
	type Message struct {
		Message   string `json:"message"`
		PublicKey string `json:"pubKey"`
		Signature string `json:"signature"`
	}
	var newMess Message
	json.Unmarshal([]byte(test), &newMess)

	bpub, _ := cryptography.ConvertBase64ToBytes(newMess.PublicKey)
	bsig, _ := cryptography.ConvertBase64ToBytes(newMess.Signature)

	ok, err := cryptography.VerifySig(newMess.Message, bsig, bpub)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)

}
