package controllers

func (server *Server) initializeRoutes() {

	server.Router.HandleFunc("/account", server.CreateAccount).Methods("POST")

}