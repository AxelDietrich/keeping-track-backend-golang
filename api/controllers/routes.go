package controllers

func (server *Server) initializeRoutes() {

	//Accounts
	server.Router.HandleFunc("/account", server.CreateAccount).Methods("POST")
	server.Router.HandleFunc("/login", server.Login).Methods("POST")

	//Subcategories
	server.Router.HandleFunc("/{categoryID}/subcategory", server.CreateSubcategory).Methods("POST")
	server.Router.HandleFunc("/subcategory/{subcategoryID}", server.ModifySubcategory).Methods("PATCH").Queries("name", "{name}")
	server.Router.HandleFunc("/subcategory/{subcategoryID}", server.DeleteSubcategory).Methods("DELETE")
	server.Router.HandleFunc("/{categoryID}/subcategories", server.GetAllSubcategories).Methods("GET")

	//Balances
	server.Router.HandleFunc("/{accountID}/balance", server.GetBalance).Methods("GET")

	//Records
	server.Router.HandleFunc("/{subcategoryID}/record", server.CreateRecord).Methods("POST")
	server.Router.HandleFunc("/record/{recordID}", server.UpdateRecord).Methods("PUT")
	server.Router.HandleFunc("/record/{recordID}", server.DeleteRecord).Methods("DELETE")
	server.Router.HandleFunc("/{subcategoryID}/records", server.GetAllDebtRecordsBySubcategoryID).Methods("GET")

	//DebtRecords
	server.Router.HandleFunc("/{subcategoryID}/record/debt", server.CreateDebtRecord).Methods("POST")
	server.Router.HandleFunc("/record/debt/{recordID}", server.DeleteDebtRecord).Methods("DELETE")
	server.Router.HandleFunc("/record/debt/{recordID}", server.UpdateDebtRecord).Methods("PUT")
	server.Router.HandleFunc("/{subcategoryID}/records/debt", server.GetAllDebtRecordsBySubcategoryID).Methods("GET")

}
