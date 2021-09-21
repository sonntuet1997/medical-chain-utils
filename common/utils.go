package common

import (
	"log"

	"github.com/google/uuid"
)

func StrPtr(a string) *string {
	return &a
}
func StrVal(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

func BoolPtr(a bool) *bool {
	return &a
}
func BoolVal(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
}

func IntPtr(a int) *int {
	return &a
}
func IntVal(a *int) int {
	if a == nil {
		return 0
	}
	return *a
}
func Int32Ptr(a int32) *int32 {
	return &a
}

func Bytes2UUID(a []byte) uuid.UUID {
	if len(a) != 16 {
		log.Println("parse uuid fail")
		return uuid.Nil
	}
	var id [16]byte
	copy(id[:], a)
	return uuid.UUID(id)
}
