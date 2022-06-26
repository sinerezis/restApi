package model

import "testing"

// Тестовый пользователь
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test@test.test",
		Password: "testtesttest",
	}

}
