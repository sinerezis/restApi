package apiserver

import (
	"io"
	"log"
	"net/http"
	"resApi/internal/app/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Структура, реализующая апи-сервер
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// Функция, создающий новый экземпляр структуры апи-сервера
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Функция запуска сервера
func (s *APIserver) Start() error {
	if err := s.configureLoger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	//при успешном старте логгер выводит сообщение
	s.logger.Info("Starting api server!")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Функция для настройки логера
// и изменения его уровня
func (s *APIserver) configureLoger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	s.logger.SetLevel(level)

	return nil
}

//функция для настройки роутера
func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

// Тестовый хэндлер
func (s *APIserver) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello, rest api!")
	}

}
