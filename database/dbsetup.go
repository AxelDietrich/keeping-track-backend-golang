package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
)

func EstablishDatabaseConnection() (*sql.DB, error) {

	db, err := sql.Open("pgx", "postgres://postgres:haliax@localhost:5432/keepingtrack")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
