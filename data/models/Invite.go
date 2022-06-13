package models

import (
	"errors"

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
}

func (i *Invite) Validate() error {
	if i.From.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing from uuid")
	}

	if i.To.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("missing to uuid")
	}

	if i.Message == "" {
		return errors.New("missing message")
	}

	return nil
}

func (i *Invite) Create(db *gorm.DB, uId, cId uuid.UUID) error {
	i.To = cId
	i.From = uId

	if err := i.Validate(); err != nil {
		return err
	}

	return db.Create(&i).Take(&i).Error
}

func (i *Invite) Update(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Invite{}).Where("id = ?", id).Updates(i).Take(&i).Error
}

func (i *Invite) FindById(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Invite{}).Where("id = ?", id).Take(&i).Error
}

func (i *Invite) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(i, id).Error
}