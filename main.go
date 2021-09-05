package main

import (
	"fmt"
	"log"

	"github.com/sonntuet1997/medical-chain-utils/cryptography"
)

func main() {
	a, err := cryptography.GenAuthorization("94e914e1-eb58-48ff-9dac-58c19c55896c", "")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("a: %v\n", a)
}
