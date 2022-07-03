package teststore

import (
	"resApi/internal/app/store"
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
func New() *Store {
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
