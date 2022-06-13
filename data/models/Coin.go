package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coin struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name                string    `gorm:"size:18; not null" json:"name"`
	ProfilePicture      string    `gorm:"size:200; not null" json:"profile_picture"`
	TokenName           string    `gorm:"size:4; not null" json:"token_name"`
	SubtokenName        string    `gorm:"size:18; not null" json:"subtoken_name"`
	DefaultBalance      uint      `gorm:"not null" json:"default_balance"`
	DefaultBalanceScale uint      `gorm:"not null" json:"default_balance_scale"`
}

func (c *Coin) Create(db *gorm.DB) (*Coin, error) {
	err := db.Create(&c).Error
	if err != nil {
		return &Coin{}, err
	}

	return c, err
}

func (c *Coin) Update(db *gorm.DB, id uuid.UUID) (*Coin, error) {
	coin := &Coin{}
	err := db.Model(&Coin{}).Where("id = ?", id).Updates(c).Take(&coin).Error
	if err != nil {
		return &Coin{}, err
	}

	return coin, err
}

func (c *Coin) FindById(db *gorm.DB, id uuid.UUID) (*Coin, error) {
	err := db.Model(&Coin{}).Where("id = ?", id).Take(&c).Error
	if err != nil {
		return &Coin{}, err
	}

	return c, err
}

func (c *Coin) FindAll(db *gorm.DB, pagination int, limit int) (*[]Coin, error) {
	coins := []Coin{}
	err := db.Debug().Model(&Coin{}).Limit(limit).Offset(pagination * limit).Find(&coins).Error
	if err != nil {
		return &[]Coin{}, err
	}

	return &coins, err
}