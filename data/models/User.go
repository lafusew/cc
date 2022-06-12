package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	DisplayName    string    `gorm:"size:18; not null" json:"display_name"`
	Tag            string    `gorm:"size:4; not null" json:"tag"`
	ProfilePicture string    `gorm:"size:200;" json:"profile_picture"`
	Desc           string    `gorm:"size:255;" json:"desc"`
}
