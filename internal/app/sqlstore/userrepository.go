package sqlstore

import (
	"database/sql"

	"github.com/vielendanke/restful-service/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	db *sql.DB
}

// FindAll ...
func (ur *UserRepository) FindAll() ([]model.User, error) {
	row, err := ur.db.Query("SELECT id, username, nickname FROM users")
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for row.Next() {
		user := &model.User{}
		row.Scan(&user.ID, &user.Username, &user.Nickname)
		users = append(users, *user)
	}
	return users, nil
}

// Save ...
func (ur *UserRepository) Save(user *model.User) error {
	row := ur.db.QueryRow(
		"INSERT INTO users(id, username, encrypted_password, nickname, authority) VALUES($1, $2, $3, $4, $5)",
		user.ID, user.Username, user.EncryptedPassword, user.Nickname, user.Authority,
	)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

// Find ...
func (ur *UserRepository) Find(id string) error {
	return nil
}

// FindByUsername ...
func (ur *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := ur.db.QueryRow(
		"SELECT id, username, encrypted_password, nickname, authority FROM users WHERE username=$1",
		username,
	).Scan(&user.ID, &user.Username, &user.EncryptedPassword, &user.Nickname, &user.Authority); err != nil {
		return nil, err
	}
	return user, nil
}

// Delete ...
func (ur *UserRepository) Delete(id string) error {
	return nil
}
