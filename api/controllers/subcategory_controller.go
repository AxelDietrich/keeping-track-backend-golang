package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (server *Server) CreateSubcategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	catID, err := strconv.Atoi(vars["categoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, _ := strconv.Atoi(r.Header.Get("Userid"))
	err = repositories.CheckIfUserCategory(server.DB, catID, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var subcategory models.Subcategory
	err = json.Unmarshal(reqBody, &subcategory)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subcategory.CategoryID, err = strconv.Atoi(vars["categoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	prepareSubcategory(&subcategory)
	err = validateSubcategory(&subcategory)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.CreateSubcategory(server.DB, &subcategory)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Subcategory successfully created")

}

func (server *Server) ModifySubcategory(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["subcategoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	accID, err := strconv.Atoi(r.Header.Get("Userid"))
	err = repositories.CheckIfUserSubcategory(server.DB, id, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	err = repositories.UpdateSubcategory(server.DB, id, vars["name"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSONString(w, http.StatusOK, "Subcategory successfully updated")

}

func (server *Server) DeleteSubcategory(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	subID, err := strconv.Atoi(vars["subcategoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	accID, err := strconv.Atoi(r.Header.Get("Userid"))
	err = repositories.CheckIfUserSubcategory(server.DB, subID, accID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	err = repositories.DeleteSubcategory(server.DB, subID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSONString(w, http.StatusOK, "Subcategory successfully deleted")
}

func (server *Server) GetAllSubcategories(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["categoryID"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Invalid categoryID"))
		return
	}
	accID, _ := strconv.Atoi(r.Header.Get("Userid"))
	err = repositories.CheckIfUserCategory(server.DB, categoryID, accID)
	subcategories, err := repositories.GetAllSubcategories(server.DB, categoryID)
	responses.JSON(w, http.StatusOK, subcategories)

}

func validateSubcategory(s *models.Subcategory) error {

	if s.Name == "" {
		return errors.New("Subcategory name is required")
	}
	return nil
}

func prepareSubcategory(s *models.Subcategory) {

	s.Amount = 0
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}
