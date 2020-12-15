package serviceimpl

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/store"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceImpl ...
type UserServiceImpl struct {
	store store.Store
}

// SaveUser ...
func (us *UserServiceImpl) SaveUser(user *model.User) error {
	var err error
	user.Authority = model.RoleUser
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

// FindByUsername ...
func (us *UserServiceImpl) FindByUsername(username string) (*model.User, error) {
	user, err := us.store.UserRepository().FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login ...
func (us *UserServiceImpl) Login(username string, password string) (*model.User, error) {
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
func (us *UserServiceImpl) FindAllUsers() ([]model.User, error) {
	return us.store.UserRepository().FindAll()
}
