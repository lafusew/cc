package models

import (
	"errors"

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

func (t *Transaction) Validate() error {
	if t.Amount == 0 {
		return errors.New("missing amount")
	}

	if t.CoinID.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing coin_id")
	}

	if t.From.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing from uuid")
	}

	if t.To.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing to uuid")
	}

	if t.Label == "" {
		return errors.New("missing label")
	}

	return nil
}

func (t *Transaction) Create(db *gorm.DB, toUId, fromUId, cId uuid.UUID) error {
	t.CoinID = cId
	t.To = toUId
	t.From = fromUId

	if err := t.Validate(); err != nil {
		return err
	}

	return db.Create(&t).Take(&t).Error
}

func (t *Transaction) Update(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Transaction{}).Where("id = ?", id).Updates(t).Take(&t).Error
}

func (t *Transaction) FindById(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Transaction{}).Where("id = ?", id).Take(&t).Error
}

func (t *Transaction) FindByToUserId(db *gorm.DB, uId uuid.UUID) error {
	return db.Model(&Transaction{}).Where("to = ?", uId).Take(&t).Error
}

func (t *Transaction) FindByFromUserId(db *gorm.DB, uId uuid.UUID) error {
	return db.Model(&Transaction{}).Where("from = ?", uId).Take(&t).Error
}

func (t *Transaction) FindByCoinId(db *gorm.DB, cId uuid.UUID) error {
	return db.Model(&Transaction{}).Where("coin_id = ?", cId).Take(&t).Error
}

func (t *Transaction) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(t, id).Error
}