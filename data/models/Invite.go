package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invite struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	From     uuid.UUID `gorm:"not null" json:"from"`
	FromUser User      `gorm:"foreignKey:From"`
	To       uuid.UUID `gorm:"not null" json:"to"`
	ToCoin   Coin      `gorm:"foreignKey:To"`
	Message  string    `gorm:"size:200; not null" json:"message"`
	ExpireAt time.Time `gorm:"not null" json:"expire_at"`
}

func (a *Invite) Create(db *gorm.DB, uId, cId uuid.UUID) error {
	a.To = cId
	a.From = uId

	return db.Create(&a).Take(&a).Error
}

func (a *Invite) Update(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Invite{}).Where("id = ?", id).Updates(a).Take(&a).Error
}

func (a *Invite) FindById(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Invite{}).Where("id = ?", id).Take(&a).Error
}

func (a *Invite) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(a, id).Error
}