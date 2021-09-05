package common

import (
	"github.com/google/uuid"
)

func StrPtr(a string) *string {
	return &a
}

func Bytes2UUID(a []byte) uuid.UUID {
	id, err := uuid.ParseBytes(a)
	if err != nil {
		return uuid.Nil
	}
	return id
}
