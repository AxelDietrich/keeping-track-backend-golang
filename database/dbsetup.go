package database

import (
	"database/sql"
)

func EstablishDatabaseConnection() (*sql.DB, error) {

	//pgURL, err := pq.ParseURL("postgres://mbdvayqu:ywLhPz6ssFEjCoIoRhYil3xxfyb8ZSmh@motty.db.elephantsql.com:5432/mbdvayqu")
	//if err != nil {
	//	return nil, err
	//}

	db, err := sql.Open("pgx", "postgres://mbdvayqu:ywLhPz6ssFEjCoIoRhYil3xxfyb8ZSmh@motty.db.elephantsql.com:5432/mbdvayqu")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	/*pgConn, err := pgconn.Connect(context.Background(), "postgres://mbdvayqu:ywLhPz6ssFEjCoIoRhYil3xxfyb8ZSmh@motty.db.elephantsql.com:5432/mbdvayqu")
	if err != nil {
		return nil, err
	}
	return pgConn, nil*/
	return db, nil
}
