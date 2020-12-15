package testservice

import (
	"github.com/vielendanke/restful-service/internal/app/model"
)

type UserServiceMock struct {}

func (um *UserServiceMock) SaveUser(user *model.User) error {
	return nil
}

// FindByUsername ...
func (um *UserServiceMock) FindByUsername(username string) (*model.User, error) {
	return model.MockUser(), nil
}

// Login ...
func (um *UserServiceMock) Login(username string, password string) (*model.User, error) {
	return model.MockUser(), nil
}

// FindAllUsers ...
func (um *UserServiceMock) FindAllUsers() ([]model.User, error) {
	return []model.User{}, nil
}
