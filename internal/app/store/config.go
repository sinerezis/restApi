package store

// Структура конфига
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// Вспомогательная функция для создания
// нового конфига
func NewConfig() *Config {
	return &Config{}
}
