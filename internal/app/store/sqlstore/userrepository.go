package sqlstore

import "resApi/internal/app/model"

// Репозиторий
type UserRepository struct {
	Store *Store
}

// Создание нового пользователя
func (r *UserRepository) Create(u *model.User) error {
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
func (r *UserRepository) FindByEmaiL(email string) (*model.User, error) {
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
		return nil, err
	}

	return &u, nil
}
