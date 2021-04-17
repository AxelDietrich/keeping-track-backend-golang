package controllers

import "net/http"

func (server *Server) initializeRoutes() {

	//Accounts
	server.Router.HandleFunc("/account", server.CreateAccount).Methods("POST")
	server.Router.HandleFunc("/login", server.Login).Methods("POST")

	//Subcategories
	server.Router.HandleFunc("/{categoryID}/subcategory", AuthMiddleware(server.CreateSubcategory)).Methods("POST")
	server.Router.HandleFunc("/subcategory/{subcategoryID}", AuthMiddleware(server.ModifySubcategory)).Methods("PATCH").Queries("name", "{name}")
	server.Router.HandleFunc("/subcategory/{subcategoryID}", AuthMiddleware(server.DeleteSubcategory)).Methods("DELETE")
	server.Router.HandleFunc("/{categoryID}/subcategories", AuthMiddleware(server.GetAllSubcategories)).Methods("GET")

	//Balances
	server.Router.HandleFunc("/{accountID}/balance", AuthMiddleware(server.GetBalance)).Methods("GET")

	//Records
	server.Router.HandleFunc("/{subcategoryID}/record", AuthMiddleware(server.CreateRecord)).Methods("POST")
	server.Router.HandleFunc("/record/{recordID}", AuthMiddleware(server.UpdateRecord)).Methods("PUT")
	server.Router.HandleFunc("/record/{recordID}", AuthMiddleware(server.DeleteRecord)).Methods("DELETE")
	server.Router.HandleFunc("/{subcategoryID}/records", AuthMiddleware(server.GetAllRecordsBySubcategoryID)).Methods("GET")

	//DebtRecords
	server.Router.HandleFunc("/{subcategoryID}/record/debt", AuthMiddleware(server.CreateDebtRecord)).Methods("POST")
	server.Router.HandleFunc("/record/debt/{recordID}", AuthMiddleware(server.DeleteDebtRecord)).Methods("DELETE")
	server.Router.HandleFunc("/record/debt/{recordID}", AuthMiddleware(server.UpdateDebtRecord)).Methods("PUT")
	server.Router.HandleFunc("/{subcategoryID}/records/debt", AuthMiddleware(server.GetAllDebtRecordsBySubcategoryID)).Methods("GET")

	server.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("././static/")))
}
