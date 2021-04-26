package model_test

import (
	"github.com/i3acsi/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUser_BeforeCreate ...
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
