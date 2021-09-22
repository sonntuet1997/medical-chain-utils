package main

import (
	"encoding/json"
	"fmt"

	"github.com/sonntuet1997/medical-chain-utils/cryptography"
)

var JS = `
{"data":{"providerId":"p1","name":"Provider A","manager":"!@#$%^&*()","mail":"abc@gmail.com","phone":"093112312","description":"description A","joinDate":1632060129,"code":"p1","_actionType":"PUT_V1-ADMIN-PROVIDER-P1","_timestamp":"2021-09-22T06:15:34.583Z"},
"_signature":"ZqHH1L/XpnWo6hFTzujDTrJuh4EeFOSlX6Av8dRMGA81T1mOwF6gHMgqPsHZP8xnCIjnC4YUeZz3YTQ+A+GLAQ=="}
`

func main() {
	var message struct {
		Data json.RawMessage `json:"data"`
	}
	json.Unmarshal([]byte(JS), &message)
	fmt.Printf("len(message.Data): %v\n", len(message.Data))
	a, _ := cryptography.ConvertMessage(message.Data)
	fmt.Printf("a: %v\n", a)

}
