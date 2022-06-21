package main

import (
	"flag"
	"log"
	"resApi/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

// Путь к файлу конфига
var (
	configPath string
)

// Эта функция позволяет нам запускать наш сервер с флагом
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	// Парсим флаг
	flag.Parse()

	// Сохраняем нвоый конфиг
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
