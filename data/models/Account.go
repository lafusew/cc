package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountAuthority string

const (
	OWNER AccountAuthority = "OWNER"
	ADMIN AccountAuthority = "ADMIN"
	USER  AccountAuthority = "USER"
)

type Account struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"size:18; not null" json:"name"`
	UserID    uuid.UUID `gorm:"not null" json:"user_id"`
	User      User
	CoinID    uuid.UUID `gorm:"not null" json:"coin_id"`
	Coin      Coin
	Authority AccountAuthority `gorm:"not null" json:"authority"`
	Balance   uint             `gorm:"not null" json:"balance"`
	Scale     uint             `gorm:"not null" json:"scale"`
}
