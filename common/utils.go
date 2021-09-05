package common

import (
	"log"

	"github.com/google/uuid"
)

func StrPtr(a string) *string {
	return &a
}

func Bytes2UUID(a []byte) uuid.UUID {
	if len(a) != 16 {
		log.Println("parse uuid fail")
	}
	var id [16]byte
	copy(id[:], a)
	return uuid.UUID(id)
}
