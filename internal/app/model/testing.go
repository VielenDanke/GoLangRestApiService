package model

import (
	"github.com/google/uuid"
	"testing"
)

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Username: "test@mail.ru",
		Password: "testuser",
		Nickname: "TestUser",
	}
}

// TestPost ...
func TestPost(t *testing.T) *Post {
	return &Post{
		Name:    "Test post name",
		Content: "Test post content about test post name",
	}
}

func MockUser() *User {
	return &User{
		ID: uuid.New().String(),
		Username: "username@mail.ru",
		Password: "userpassword",
		Nickname: "User",
		Authority: RoleUser,
	}
}

func MockPost() *Post {
	return &Post{
		ID: uuid.New().String(),
		Name:    "Test post name",
		Content: "Test post content about test post name",
	}
}
