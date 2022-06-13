package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coin struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name                string    `gorm:"size:18; not null" json:"name"`
	ProfilePicture      string    `gorm:"size:200;" json:"profile_picture"`
	TokenName           string    `gorm:"size:4; not null" json:"token_name"`
	SubtokenName        string    `gorm:"size:18; not null" json:"subtoken_name"`
	DefaultBalance      uint      `gorm:"not null" json:"default_balance"`
}



func (c *Coin) Validate() error {
	if c.Name == "" {
		return errors.New("missing name")
	}

	if c.TokenName == "" {
		return errors.New("missing token_name")
	}

	if c.SubtokenName == "" {
		return errors.New("missing subtoken_name")
	}

	if c.DefaultBalance == 0 {
		return errors.New("missing default_balance")
	}
	return nil
}

func (c *Coin) Create(db *gorm.DB) error {
	if err := c.Validate(); err != nil {
		return err
	}

	return db.Create(&c).Take(&c).Error
}

func (c *Coin) Update(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Coin{}).Where("id = ?", id).Updates(c).Take(&c).Error
}

func (c *Coin) FindById(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Coin{}).Where("id = ?", id).Take(&c).Error
}

func (c *Coin) FindAll(db *gorm.DB, pagination int, limit int) (*[]Coin, error) {
	coins := []Coin{}
	err := db.Debug().Model(&Coin{}).Limit(limit).Offset(pagination * limit).Find(&coins).Error

	return &coins, err
}

func (c *Coin) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(c, id).Error
}