package models

import "gorm.io/gorm"

type AuthType string

const (
	BASIC      AuthType = "BASIC"
	MAGIC_LINK AuthType = "MAGIC_LINK"
)

type Auth struct {
	gorm.Model
	Username string   `gorm:"size:50; not null" json:"username"`
	Password string   `gorm:"size:50; not null" json:"password"`
	Type     AuthType `gorm:"not null" json:"auth_type"`
	UserID   uint     `gorm:"not null" json:"user_id"`
}
