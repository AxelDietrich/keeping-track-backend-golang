package database

import (
	"database/sql"
)

func EstablishDatabaseConnection() (*sql.DB, error) {

	db, err := sql.Open("pgx", "postgres://mbdvayqu:ywLhPz6ssFEjCoIoRhYil3xxfyb8ZSmh@motty.db.elephantsql.com:5432/mbdvayqu")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
