package sqlstore

import (
	"database/sql"
	"resApi/internal/app/store"

	// Анонимный импорт
	_ "github.com/lib/pq"
)

// Структура для БД
type Store struct {

	// Непосредственно, сама БД
	db *sql.DB

	// Позволяет обращаться к репозиторию
	// не только из хранилища
	UserRepository *UserRepository
}

// Метод для создания нового экземпляра Store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Store.User.Create()
func (s *Store) User() store.UserRepository {

	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		Store: s,
	}
	return s.UserRepository

}
