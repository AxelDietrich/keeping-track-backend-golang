package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/requests"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"strconv"
)

func (server *Server) MoveFundsToSavings(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	amount := &requests.MoveToSavingsRequest{}
	err = json.Unmarshal(reqBody, amount)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, err := strconv.Atoi(vars["accountID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	err = repositories.MoveFundsToSavings(server.DB, amount.Amount, accID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Funds successfully moved to savings")

}
