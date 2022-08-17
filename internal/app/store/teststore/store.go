package teststore

import (
	"resApi/internal/app/model"
	"resApi/internal/app/store"
)

// Структура для БД
type Store struct {

	// Позволяет обращаться к репозиторию
	// не только из хранилища
	userRepository *UserRepository
}

// Метод для создания нового экземпляра Store
func New() *Store {
	return &Store{}
}

//Store.User.Create()
func (s *Store) User() store.UserRepository {

	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository

}
