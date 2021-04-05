package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"strconv"
	"time"
)

func (server *Server) CreateSubcategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	/*if len(vars) == 0 {
		responses.ERROR(w, 400, errors)
	}*/
	reqBody, _ := ioutil.ReadAll(r.Body)
	var subcategory models.Subcategory
	err = json.Unmarshal(reqBody, &subcategory)
	if err != nil {
		responses.ERROR(w, 400, err)
		return
	}
	subcategory.CategoryID, err = strconv.Atoi(vars["categoryID"])
	if err != nil {
		responses.ERROR(w, 400, err)
		return
	}
	prepareSubcategory(&subcategory)
	err = validateSubcategory(&subcategory)
	if err != nil {
		responses.ERROR(w, 400, err)
		return
	}
	err = repositories.CreateSubcategory(server.DB, &subcategory)
	if err != nil {
		responses.ERROR(w, 400, err)
		return
	}
	responses.JSONString(w, 200, "Subcategory successfully created")

}

func (server *Server) ModifySubcategory(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["subcategoryID"])
	err = repositories.UpdateSubcategory(server.DB, id, vars["name"])
	if err != nil {
		responses.ERROR(w, 400, err)
		return
	}
	responses.JSONString(w, 200, "Subcategory successfully updated")

}

func (server *Server) GetAllSubcategories(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["categoryID"])
	if err != nil {
		responses.ERROR(w, 400, errors.New("Invalid categoryID"))
		return
	}
	subcategories, err := repositories.GetAllSubcategories(server.DB, categoryID)
	responses.JSON(w, 200, subcategories)

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
