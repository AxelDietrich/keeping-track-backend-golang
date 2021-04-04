package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	m "keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"net/http"
	"regexp"
	"strings"
	"errors"
)

func (server *Server) CreateAccount (w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account m.Account
	err = json.Unmarshal(reqBody, &account)
	if err != nil {
		return
	}
	err = validateAccount("", &account)
	if err != nil {
		return
	}
	repositories.CreateAccount(server.DB, &account)

	fmt.Println("Endpoint Hit: Create Account")
	json.NewEncoder(w).Encode(account)
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
		if a.Email == "" {
			return errors.New("Required Email")
		}
		if !re.MatchString(a.Email) {
			return errors.New("Invalid email format")
		}
		return nil
	}
}