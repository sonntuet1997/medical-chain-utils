package utils

import (
	"github.com/google/uuid"
)

type ModelBase struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	UpdatedAt int64     `gorm:"autoUpdateTime" json:"updatedAt"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt int64     `gorm:"autoDeleteTime" json:"deletedAt"`
}
