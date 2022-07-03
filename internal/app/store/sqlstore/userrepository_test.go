package sqlstore_test

import (
	"resApi/internal/app/model"
	"resApi/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)
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
