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
		return
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
	accID, _ := strconv.Atoi(r.Header.Get("userID"))
	err = repositories.CheckIfUserRecord(server.DB, recID, accID)
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
	accID, _ := strconv.Atoi(r.Header.Get("userID"))
	err = repositories.CheckIfUserRecord(server.DB, recID, accID)
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

func (server *Server) GetAllRecordsBySubcategoryID(w http.ResponseWriter, r *http.Request) {

	var err error
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
		return
	}
	records, err := repositories.GetAllRecordsBySubcategoryID(server.DB, subID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, records)
}
