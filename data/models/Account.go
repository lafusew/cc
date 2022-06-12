package data

import "gorm.io/gorm"

type AccountAuthority string

const (
	OWNER AccountAuthority = "OWNER"
	ADMIN AccountAuthority = "ADMIN"
	USER AccountAuthority = "USER"
)

type Account struct {
	gorm.Model
	Name string `gorm:"size:18; not null" json:"name"`
	UserID uint `gorm:"not null" json:"user_id"`
	CoinID uint `gorm:"not null" json:"coin_id"`
	Authority AccountAuthority `gorm:"not null" json:"authority"`
	Balance uint `gorm:"not null" json:"balance"`
	Scale uint `gorm:"not null" json:"scale"`
}