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
	Authority AccountAuthority `gorm:"type:account_authority; not null" json:"authority"`
	Balance   uint             `gorm:"not null" json:"balance"`
	Scale     uint             `gorm:"not null" json:"scale"`
}

func (a *Account) Create(db *gorm.DB, uId, cId uuid.UUID) error {
	a.CoinID = cId
	a.UserID = uId

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