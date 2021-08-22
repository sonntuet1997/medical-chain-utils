package utils

import "time"

type ModelBase struct {
	ID       string    `gorm:"primaryKey;column:id" json:"id"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"updateAt"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"createAt"`
}

type Model interface {
	GetName() string
}
