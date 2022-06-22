package store

import (
	"database/sql"

	// Анонимный импорт
	_ "github.com/lib/pq"
)

// Структура для БД
type Store struct {

	// Файл конфига нашей БД
	config *Config

	// Непосредственно, сама БД
	db *sql.DB

	// Позволяет обращаться к репозиторию
	// не только из хранилища
	UserRepository *UserRepository
}

// Метод для создания нового экземпляра Store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Метод для подклчения к БД
func (s *Store) Open() error {

	// Подключаемся к БД
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	// Пингуем БД
	if err := db.Ping(); err != nil {
		return err
	}

	// Если все окей - записываем переммуню БД
	// в структуру Store
	s.db = db

	return nil
}

// Метод для отключения от БД (при выключении сервера, и тд )
func (s *Store) Close() {

	// Закрываем соединение
	s.db.Close()
}

// Метод необходим для того, что бы можно было использовать
// репозиторий  не только из хранилища
//
//Store.User.Create()
func (s *Store) User() *UserRepository {

	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		Store: s,
	}
	return s.UserRepository

}
