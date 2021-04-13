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

func (server *Server) CreateRecord(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	rec := &requests.AddRecordRequest{}
	err = json.Unmarshal(reqBody, rec)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	subID, err := strconv.Atoi(vars["subcategoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	err = repositories.CreateRecord(server.DB, rec, subID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Record succesfully created")
}

func (server *Server) UpdateRecord(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	recordRequest := &requests.AddRecordRequest{}
	err = json.Unmarshal(reqBody, recordRequest)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	recID, err := strconv.Atoi(vars["recordID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.UpdateRecord(server.DB, recID, recordRequest)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "The record has been successfully updated")

}

func (server *Server) DeleteRecord(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	recID, err := strconv.Atoi(vars["recordID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.DeleteRecord(server.DB, recID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Record successfully deleted")
}
