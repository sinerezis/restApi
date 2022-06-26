package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// Структура - модель данных для нашей таблицы Users
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// Валидация данных пользователя
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(8, 100)),
	)

}

// Перед сохранением пользователя пробуем
// захешировать его пароль и сохранить результат
// в модель конкретного пользователя
func (u *User) BeforeCreate() error {

	if err := u.Validate(); err != nil {
		return err
	}

	if len(u.Password) > 0 {
		enc, err := encryptedString(u.Password)
		if err != nil {
			return nil
		}

		u.EncryptedPassword = enc
	}
	return nil
}

// Функция хэширования пароля пользователя
func encryptedString(s string) (string, error) {

	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
