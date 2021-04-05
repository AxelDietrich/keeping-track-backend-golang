package controllers

func (server *Server) initializeRoutes() {

	server.Router.HandleFunc("/account", server.CreateAccount).Methods("POST")
	server.Router.HandleFunc("/login", server.Login).Methods("POST")
	server.Router.HandleFunc("/{categoryID}/subcategory", server.CreateSubcategory).Methods("POST")
	server.Router.HandleFunc("/{categoryID}/subcategories", server.GetAllSubcategories).Methods("GET")
	server.Router.HandleFunc("/subcategory/{subcategoryID}", server.ModifySubcategory).Methods("PATCH").Queries("name", "{name}")
	server.Router.HandleFunc("/{subcategoryID}/record", server.CreateRecord).Methods("POST")
	server.Router.HandleFunc("/{accountID}/savings", server.MoveFundsToSavings).Methods("POST")
	server.Router.HandleFunc("/subcategory/{subcategoryID}", server.DeleteSubcategory).Methods("DELETE")
}
