package data

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name string `gorm:"size:18; not null" json:"name"`
	UserID uint ``
	CoinID uint
	Authority string 
	Balance uint
	Scale uint
}