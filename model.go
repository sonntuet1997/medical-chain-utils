package utils

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"updateAt"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"createAt"`
	DeleteAt time.Time `gorm:"autoDeleteTime" json:"deleteAt"`
}
