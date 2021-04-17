package controllers

import (
	"encoding/json"
	"io/ioutil"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/requests"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) MoveFundsToSavings(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	amount := &requests.AmountRequest{}
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

func (server *Server) AddIncome(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	amount := &requests.AmountRequest{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, amount)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, err := strconv.Atoi(vars["accountID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.AddIncome(server.DB, amount.Amount, accID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Funds succesfully added")

}

func (server *Server) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accID, err := strconv.Atoi(vars["accountID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	b := &models.Balance{}
	b, err = repositories.GetBalance(server.DB, accID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, b)
}
