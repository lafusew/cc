package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Label  string `gorm:"size:100; not null" json:"label"`
	Amount uint   `gorm:"not null" json:"amount"`
	Scale  uint   `gorm:"not null" json:"scale"`
	CoinID uint   `gorm:"not null" json:"coin_id"`
	From   uint   `gorm:"not null" json:"from"`
	To     uint   `gorm:"not null" json:"to"`
}
