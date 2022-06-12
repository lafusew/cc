package models

import (
	"time"

	"gorm.io/gorm"
)

type Invite struct {
	gorm.Model
	From     uint      `gorm:"not null" json:"by"`
	To       uint      `gorm:"not null" json:"from"`
	Message  string    `gorm:"size:200; not null" json:"message"`
	ExpireAt time.Time `gorm:"not null" json:"expire_at"`
}
