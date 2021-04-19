package controllers

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	m "keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/requests"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"regexp"
	"strings"
	"time"

)

func (server *Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	signUp := &requests.SignUp{}
	err = json.Unmarshal(reqBody, &signUp)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	account := &m.Account{}
	account.Email = signUp.Email
	account.Username = signUp.Username
	account.Password = signUp.Password
	err = validateAccount("", account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	BeforeSave(account)

	accountPersisted, err := repositories.CreateAccount(server.DB, account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	} else {
		responses.JSON(w, http.StatusOK, accountPersisted)
	}
}

func BeforeSave(a *m.Account) error {
	a.Username = html.EscapeString(strings.TrimSpace(a.Username))
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.Email = strings.ToLower(a.Email)
	a.CreatedAt = time.Now()
	hashedPassword, err := m.Hash(a.Password)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
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
