package store

// Структура для БД
type Store struct {
	config *Config
}

// Метод для создания нового экземпляра Store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Метод для подклчения к БД
func (s *Store) Open() error {
	return nil
}

// Метод для отключения от БД (при выключении сервера, и тд )
func (s *Store) Close() {

}
