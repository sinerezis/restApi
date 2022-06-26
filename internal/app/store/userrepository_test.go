package store_test

import (
	"resApi/internal/app/model"
	"resApi/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "test@test.test"
	_, err := s.User().FindByEmaiL(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(&model.User{
		Email: "test@test.test",
	})

	u, err = s.User().FindByEmaiL(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
