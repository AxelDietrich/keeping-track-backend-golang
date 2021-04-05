package controllers

func (server *Server) initializeRoutes() {

	server.Router.HandleFunc("/account", server.CreateAccount).Methods("POST")
	server.Router.HandleFunc("/login", server.Login).Methods("POST")

}
