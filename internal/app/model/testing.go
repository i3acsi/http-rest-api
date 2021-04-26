package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

// TestUserFindByEmail ...
func TestUserFindByEmail(t *testing.T) *User {
	return &User{
		Email:    "user@findByEmail.org",
		Password: "password",
	}
}
