package store_test

import (
	"github.com/i3acsi/http-rest-api/internal/app/model"
	"github.com/i3acsi/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u := model.TestUserFindByEmail(t)
	email := u.Email
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	s.User().Create(model.TestUserFindByEmail(t))
	usr, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, usr)
}
