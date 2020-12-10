package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	// RoleUser ...
	RoleUser = iota + 1
	// RoleAdmin ...
	RoleAdmin
)

// User ...
type User struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Password          string `json:"-"`
	Nickname          string `json:"nickname"`
	EncryptedPassword string `json:"-"`
	Authority         int    `json:"-"`
}

// BeforeSaving ...
func (u *User) BeforeSaving() error {
	u.ID = uuid.New().String()
	pass, err := encryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = pass
	u.Password = ""
	return nil
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Username, validation.Required, is.Email),
		validation.Field(&u.Nickname, validation.Required, validation.Length(3, 50)),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func encryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
