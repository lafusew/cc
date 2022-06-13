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

func (t *Transaction) Create(db *gorm.DB, toUId, fromUId, cId uuid.UUID) error {
	t.CoinID = cId
	t.To = toUId
	t.From = fromUId

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