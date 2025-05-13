package infrastracture

import (
	"database/sql"
	"time"
)

func NewPostDB() (*sql.DB, error) {

	connectionString := "postgresql://chera_mihiretu:123456@localhost:5432/ethiopia_market_of_digitals?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
