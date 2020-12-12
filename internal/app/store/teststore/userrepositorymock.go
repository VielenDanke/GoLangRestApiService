package teststore

import (
	"fmt"

	"github.com/vielendanke/restful-service/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	userDB map[string]model.User
}

// Save ...
func (ur *UserRepository) Save(user *model.User) error {
	ur.userDB[user.ID] = *user
	return nil
}

// Find ...
func (ur *UserRepository) Find(id string) (*model.User, error) {
	user, ok := ur.userDB[id]
	if !ok {
		return nil, fmt.Errorf("User not found")
	}
	return &user, nil
}

// FindByUsername ...
func (ur *UserRepository) FindByUsername(username string) (*model.User, error) {
	for _, v := range ur.userDB {
		if v.Username == username {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

// Delete ...
func (ur *UserRepository) Delete(id string) error {
	delete(ur.userDB, id)
	return nil
}

// FindAll ...
func (ur *UserRepository) FindAll() ([]model.User, error) {
	users := []model.User{}
	for _, v := range ur.userDB {
		users = append(users, v)
	}
	return users, nil
}
