package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	m "keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"regexp"
	"strings"
)

func (server *Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account m.Account
	err = json.Unmarshal(reqBody, &account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = validateAccount("", &account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	accountPersisted, err := repositories.CreateAccount(server.DB, &account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	} else {
		responses.JSON(w, http.StatusOK, accountPersisted)
	}
}

func validateAccount(action string, a *m.Account) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	switch strings.ToLower(action) {
	case "login":
		if a.Password == "" {
			return errors.New("Required Password")
		}
		if a.Email == "" {
			return errors.New("Required Email")
		}
		if !re.MatchString(a.Email) {
			return errors.New("Invalid email format")
		}
		return nil

	default:
		if a.Username == "" {
			return errors.New("Required Username")
		}
		if a.Password == "" {
			return errors.New("Required Password")
		}
		if len(a.Password) < 6 {
			return errors.New("Password must have at least 6 characters")
		}
		if a.Email == "" {
			return errors.New("Required Email")
		}
		if !re.MatchString(a.Email) {
			return errors.New("Invalid email format")
		}
		return nil
	}
}
