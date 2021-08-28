package utils

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano" json:"updatedAt"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano" json:"createdAt"`
	DeletedAt time.Time `gorm:"autoDeleteTime:nano" json:"deletedAt"`
}
