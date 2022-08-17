package apiserver

// Cтруктура конфига
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

// Функция возвращает новый экземпляр конфига с параметрами,
// установленными заранее
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
