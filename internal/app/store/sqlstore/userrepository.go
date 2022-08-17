package sqlstore

import (
	"database/sql"
	"resApi/internal/app/model"
	"resApi/internal/app/store"
)

// Репозиторий
type UserRepository struct {
	Store *Store
}

// Создание нового пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err

	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.Store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id ",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// Поиск пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := model.User{
		Email: email,
	}
	if err := r.Store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound

		}
		return nil, err
	}

	return &u, nil
}
