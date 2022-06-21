package apiserver

import "resApi/internal/app/store"

// Cтруктура конфига
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// Функция возвращает новый экземпляр конфига с параметрами,
// установленными заранее
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
