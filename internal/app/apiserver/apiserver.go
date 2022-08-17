package apiserver

import (
	"database/sql"
	"net/http"
	"resApi/internal/app/store/sqlstore"
)

// Старт сервера
func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(DatabaseURL string) (*sql.DB, error) {

	db, err := sql.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
