package service

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/sqlstore"
	"golang.org/x/crypto/bcrypt"
)

// UserService ...
type UserService struct {
	store *sqlstore.Store
}

// SaveUser ...
func (us *UserService) SaveUser(user *model.User) error {
	var err error
	if err = user.BeforeSaving(); err != nil {
		return err
	}
	if err = user.Validate(); err != nil {
		return err
	}
	if err = us.store.UserRepository().Save(user); err != nil {
		return err
	}
	return nil
}

// Login ...
func (us *UserService) Login(username string, password string) (*model.User, error) {
	user, err := us.store.UserRepository().FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.EncryptedPassword), []byte(password),
	); err != nil {
		return nil, err
	}
	return user, nil
}

// FindAllUsers ...
func (us *UserService) FindAllUsers() ([]model.User, error) {
	return us.store.UserRepository().FindAll()
}
