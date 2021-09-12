package main

import (
	"fmt"
	"log"

	"github.com/sonntuet1997/medical-chain-utils/cryptography"
)

func main() {
	priv, _, err := cryptography.GenerateKeyPair(nil)
	if err != nil {
		log.Println(err)
	}
	// data := map[string]interface{}{"id": "HTpLb7OwuuAXTsVdGRmvn"}
	sig, err := cryptography.SignMessage("HTpLb7OwuuAXTsVdGRmvn", priv)
	if err != nil {
		log.Println(err)
	}
	bsig := cryptography.ConvertBytesToBase64(sig)
	str := "HTpLb7OwuuAXTsVdGRmvn" + "," + bsig
	fmt.Printf("[]byte(str): %v\n", []byte(str))
	fmt.Printf("len([]byte(str)): %v\n", len([]byte(str)))
}
