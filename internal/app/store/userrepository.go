package store

import "github.com/vielendanke/restful-service/internal/app/model"

// UserRepository ...
type UserRepository interface {
	FindAll() ([]model.User, error)
	Save(*model.User) error
	Find(string) (*model.User, error)
	FindByUsername(string) (*model.User, error)
	Delete(string) error
}
