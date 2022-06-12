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

func (u *User) Create(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, err
}

func (u *User) Update(db *gorm.DB, id uuid.UUID) (*User, error) {
	user := &User{}
	db.Model(&User{}).Where("id = ?", id).Updates(u).Take(&user)

	return user, nil
}

func (u *User) FindById(db *gorm.DB, id uuid.UUID) (*User, error) {
	err := db.Model(User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, err
}

func (u *User) FindAll(db *gorm.DB, pagination int, limit int) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(limit).Offset(pagination * limit).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}