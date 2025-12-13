package db

import (
	"bugounty/internals/config"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresDB(config config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.Host, config.Port, config.User, config.Password, config.DBName))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
