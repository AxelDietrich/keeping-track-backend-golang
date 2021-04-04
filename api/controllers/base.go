package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"keeping-track-backend-golang/database"
	"log"
	"net/http"
)

type Server struct {
	DB	*gorm.DB
	Router *mux.Router
}

var err error

func (server *Server) Initialize() error {

	server.DB, err = database.EstablishDatabaseConnection()
	if err != nil {
		return errors.New("Error al conectar")
	}
	//server.DB.Debug().AutoMigrate()
	server.Router = mux.NewRouter()

	server.initializeRoutes()
	return nil
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
