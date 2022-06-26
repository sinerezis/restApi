package model_test

import (
	"resApi/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validation(t *testing.T) {

	// Структура тестовых кейсов
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			// Корректный пользователь
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			// Пустой емейл
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			// некорректный емейл
			name: "unvalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "12341234@123.123"
				return u
			},
			isValid: false,
		},
		{
			// Пустой пароль
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			// Короткий емейл
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "1"
				return u
			},
			isValid: false,
		},
		{
			// Есть зашифрованый пароль
			name: "with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "testEncryptedPassword"
				return u
			},
			isValid: true,
		},
	}

	// Проходимся циклом по всем
	// тест кейсам
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Если валидно - то мы не ожидаем ошибку
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
				// Если не валидно - то ожидаем ошибку
			} else {
				assert.Error(t, tc.u().Validate())

			}
		})
	}

}

func TestUser_BeforeCreate(t *testing.T) {

	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
