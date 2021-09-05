package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sonntuet1997/medical-chain-utils/common"
)

func main() {
	test := uuid.New()
	a := common.Bytes2UUID(test[:]).String()
	fmt.Printf("a: %v\n", a)
}
