package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func EstablishDatabaseConnection () (*gorm.DB, error) {
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
