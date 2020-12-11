package model

import "testing"

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
