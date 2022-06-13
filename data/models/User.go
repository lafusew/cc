package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	DisplayName    string    `gorm:"size:18; not null" json:"display_name"`
	Tag            string    `gorm:"size:4; not null" json:"tag"`
	ProfilePicture string    `gorm:"size:200;" json:"profile_picture"`
	Desc           string    `gorm:"size:255;" json:"desc"`
}

func (u *User) Create(db *gorm.DB) error {
	err := db.Create(&u).Take(&u).Error
	if err != nil {
		return err
	}

	return err
}

func (u *User) Update(db *gorm.DB, id uuid.UUID) error {
	err := db.Model(&User{}).Where("id = ?", id).Updates(u).Take(&u).Error
	if err != nil {
		return err
	}

	return err
}

func (u *User) FindById(db *gorm.DB, id uuid.UUID) error {
	err := db.Model(User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return err
	}

	return err
}

func (u *User) FindAll(db *gorm.DB, us *[]User, pagination int, limit int) error {
	err := db.Debug().Model(&User{}).Limit(limit).Offset(pagination * limit).Find(&us).Error
	if err != nil {
		return err
	}

	return err
}

func (u *User) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(u, id).Error
}
