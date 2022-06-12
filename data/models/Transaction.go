package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Label    string    `gorm:"size:100; not null" json:"label"`
	Amount   uint      `gorm:"not null" json:"amount"`
	Scale    uint      `gorm:"not null" json:"scale"`
	CoinID   uuid.UUID `gorm:"not null" json:"coin_id"`
	Coin     Coin
	From     uuid.UUID `gorm:"not null" json:"from"`
	FromUser User      `gorm:"foreignKey:From"`
	To       uuid.UUID `gorm:"not null" json:"to"`
	ToUser   User      `gorm:"foreignKey:To"`
}
