package main

import (
	"encoding/json"
	"fmt"

	"github.com/sonntuet1997/medical-chain-utils/cryptography"
)

func main() {
	pk := "A0IAZ/+cyKCNhxYELQAraHhIkI+T+WWaJLpcYjMK0a+s"
	js := `{"data":{"providerId":"dsadada","name":"32131231","manager":"@#@#!##$%^&*()","mail":"3dajfajf@gmail.com","phone":"@#$%^&*()","description":"@#$%^&*()","_actionType":"POST_API-ADMIN-PROVIDER","_timestamp":"2021-09-21T13:23:47.628Z"},"_signature":"9iReEsP9LquB24Ysy0tWD3DczDoDk6MpKxfQvAYa4lZGgspisGzgB2ZAoHqd2DhsIsLE+kPN3nx2dW974yqT6Q=="}`
	type base struct {
		Data      json.RawMessage `json:"data"`
		Signature string          `json:"_signature"`
	}
	var a base
	json.Unmarshal([]byte(js), &a)
	bpk, _ := cryptography.ConvertBase64ToBytes(pk)
	bsig, _ := cryptography.ConvertBase64ToBytes(a.Signature)
	fmt.Printf("bpk: %v\n", bpk)
	fmt.Printf("bsig: %v\n", bsig)
}
