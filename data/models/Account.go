package models

import (
	"errors"

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
	Authority AccountAuthority `gorm:"type:account_authority; not null" json:"authority"`
	Balance   uint             `gorm:"not null" json:"balance"`
	Scale     uint             `gorm:"not null" json:"scale"`
}

func (a *Account) Validate() error {
	if a.CoinID.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing coin_id")
	}

	if a.UserID.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing user_id")
	}

	if a.Authority == "" {
		return errors.New("missing account_authority")
	}

	if a.Name == "" {
		return errors.New("missing name")
	}

	return nil
}

func (a *Account) Create(db *gorm.DB, uId, cId uuid.UUID) error {
	a.CoinID = cId
	a.UserID = uId

	if err := a.Validate(); err != nil {
		return err
	}

	return db.Create(&a).Take(&a).Error
}

func (a *Account) Update(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Account{}).Where("id = ?", id).Updates(a).Take(&a).Error
}

func (a *Account) FindById(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Account{}).Where("id = ?", id).Take(&a).Error
}

func (a *Account) FindByUserId(db *gorm.DB, uId uuid.UUID) error {
	return db.Model(&Account{}).Where("user_id = ?", uId).Take(&a).Error
}

func (a *Account) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(a, id).Error
}