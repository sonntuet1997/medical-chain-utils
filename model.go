package utils

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"createdAt"`
	DeletedAt time.Time `gorm:"autoDeleteTime:milli" json:"deletedAt"`
}
