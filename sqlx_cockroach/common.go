package sqlxcockroach

import (
	"time"

	"github.com/google/uuid"
)

func GetDefaultID() uuid.UUID {
	return uuid.New()
}

func GetCurrentTimeMilli() int64 {
	return time.Now().UnixMilli()
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}

func GetCurrentTimeMicro() int64 {
	return time.Now().UnixMicro()
}

func GetCurrentTimeNano() int64 {
	return time.Now().UnixNano()
}
