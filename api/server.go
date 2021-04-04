package api

import "keeping-track-backend-golang/api/controllers"

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}
