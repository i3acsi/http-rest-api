package model_test

import (
	"github.com/i3acsi/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUser_Validate ...
func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "valid with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.EncryptedPassword = "encryptedpassword"
				u.Password = ""
				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "email.ru"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "12345"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}

	u := model.TestUser(t)
	assert.NoError(t, u.Validate())
}

// TestUser_BeforeCreate ...
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
