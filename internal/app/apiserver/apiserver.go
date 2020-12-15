package apiserver

import (
	"database/sql"
	"github.com/vielendanke/restful-service/internal/app/service/serviceimpl"
	"net/http"

	"github.com/vielendanke/restful-service/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	store := sqlstore.NewStore(db)
	service := serviceimpl.NewService(store)
	server, err := NewServer(service, config)

	if err != nil {
		return err
	}
	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
