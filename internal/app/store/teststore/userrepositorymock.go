package teststore

import "github.com/vielendanke/restful-service/internal/app/model"

// UserRepository ...
type UserRepository struct {
	userDB map[string]*model.User
}

// Save ...
func (ur *UserRepository) Save(*model.User) error {
	return nil
}

// Find ...
func (ur *UserRepository) Find(string) (*model.User, error) {
	return nil, nil
}

// FindByUsername ...
func (ur *UserRepository) FindByUsername(string) (*model.User, error) {
	return nil, nil
}

// Delete ...
func (ur *UserRepository) Delete(string) error {
	return nil
}

// FindAll ...
func (ur *UserRepository) FindAll() ([]model.User, error) {
	return nil, nil
}
