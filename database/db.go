package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	User     string
	Password string
	DBName   string
	DBPort   string
	Host     string
	SSLMode  string
}

func NewPostgreSQLStorage(cfg DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.DBPort, cfg.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("DB: Successfully connected")
	return db, nil
}
