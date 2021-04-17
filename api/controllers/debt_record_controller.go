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

func (server *Server) CreateDebtRecord(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	rec := &requests.AddRecordRequest{}
	err = json.Unmarshal(reqBody, rec)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subID, err := strconv.Atoi(vars["subcategoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, _ := strconv.Atoi(r.Header.Get("userID"))
	err = repositories.CheckIfUserSubcategory(server.DB, subID, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	err = repositories.CreateDebtRecord(server.DB, rec, subID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Debt record created succesffully")

}

func (server *Server) DeleteDebtRecord(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	recID, err := strconv.Atoi(vars["recordID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, _ := strconv.Atoi(r.Header.Get("userID"))
	err = repositories.CheckIfUserDebtRecord(server.DB, recID, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.DeleteDebtRecord(server.DB, recID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Debt record succesfully deleted.")
}

func (server *Server) UpdateDebtRecord(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	req := &requests.AddRecordRequest{}
	err = json.Unmarshal(reqBody, req)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	recID, err := strconv.Atoi(vars["recordID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, _ := strconv.Atoi(r.Header.Get("userID"))
	err = repositories.CheckIfUserDebtRecord(server.DB, recID, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.UpdateDebRecord(server.DB, recID, req)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Record successfully updated")
}

func (server *Server) GetAllDebtRecordsBySubcategoryID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subID, err := strconv.Atoi(vars["subcategoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, _ := strconv.Atoi(r.Header.Get("userID"))
	err = repositories.CheckIfUserSubcategory(server.DB, subID, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	records, err := repositories.GetAllDebtRecordsBySubcategoryID(server.DB, subID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, records)
}
