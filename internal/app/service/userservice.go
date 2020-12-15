package service

import "github.com/vielendanke/restful-service/internal/app/model"

type UserService interface {
	SaveUser(user *model.User) error
	FindByUsername(string) (*model.User, error)
	Login(string, string) (*model.User, error)
	FindAllUsers() ([]model.User, error)
}
