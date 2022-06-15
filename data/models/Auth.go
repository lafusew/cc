package models

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthType string

const (
	BASIC      AuthType = "BASIC"
	MAGIC_LINK AuthType = "MAGIC_LINK"
)

type Auth struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Identifier string    `gorm:"size:50; not null" json:"identifier"`
	Password string    `gorm:"size:255; not null" json:"password"`
	Type     AuthType  `gorm:"type:auth_type; not null;default:'BASIC'" json:"auth_type"`
	UserID   uuid.UUID `gorm:"not null" json:"user_id"`
	User     User
}

func (a *Auth) HashPassword()  error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}


	a.Password = string(hashed)

	return nil
}

func (a *Auth) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}

func (a *Auth) Validate() error {
	if a.Identifier == "" {
		return errors.New("missing username")
	}

	if a.Password == "" {
		return errors.New("missing password")
	}

	if a.Type == MAGIC_LINK {
		return errors.New("unsupported login method")
	}

	return nil
}

func (a *Auth) Create(db *gorm.DB, uID uuid.UUID) error {
	a.UserID = uID

	if err := a.Validate(); err != nil {
		return err
	}

	if err := a.HashPassword(); err != nil {
		return err
	}

	return db.Create(&a).Take(&a).Error
}

func (a *Auth) Update(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Auth{}).Where("id = ?", id).Updates(a).Take(&a).Error
}

func (a *Auth) FindById(db *gorm.DB, id uuid.UUID) error {
	return db.Model(&Auth{}).Where("id = ?", id).Take(&a).Error
}

func (a *Auth) FindByIdentifier(db *gorm.DB, identifier string) error {
	return db.Model(&Auth{}).Where("identifier = ?", identifier).Take(&a).Error
}

func (a *Auth) Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(a, id).Error
}
