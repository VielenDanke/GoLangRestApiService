package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Username: "test@mail.ru",
		Password: "testuser",
		Nicname: "TestUser",
	}
}