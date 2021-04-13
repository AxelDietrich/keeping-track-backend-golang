package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func EstablishDatabaseConnectionGorm() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	dsn := "host=motty.db.elephantsql.com user=mbdvayqu password=ywLhPz6ssFEjCoIoRhYil3xxfyb8ZSmh dbname=mbdvayqu port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection failed to open")
		return nil, err
	} else {
		fmt.Println("Connection established")
		return db, nil
	}
}

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
