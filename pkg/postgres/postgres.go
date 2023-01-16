// Package postgres Функционал для работы с PostgreSQL
package postgres

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// ConnectDatabase Инициализация базы данных
func ConnectDatabase(dsn string) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile("pkg/postgres/migrations/init.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(bytes))

	if err != nil {
		return db, nil
	}

	return db, nil
}
