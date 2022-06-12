package data

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Name string `gorm:"size:18; not null" json:"name"`
	ProfilePicture string `gorm:"size:200; not null" json:"profile_picture"`
	TokenName string `gorm:"size:4; not null" json:"token_name"`
	SubtokenName string `gorm:"size:18; not null" json:"subtoken_name"`
	DefaultBalance uint `gorm:"not null" json:"default_balance"`
	DefaultBalanceScale uint `gorm:"not null" json:"default_balance_scale"`
}