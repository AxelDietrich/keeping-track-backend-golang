package controllers

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"html"
	"io/ioutil"
	m "keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"strings"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	var login m.Account
	err = json.Unmarshal(reqBody, &login)
	if err != nil {
		responses.ERROR(w, 500, err)
		return
	}
	PrepareLogin(&login)
	loginDB, err := repositories.GetAccount(server.DB, &login)
	if err != nil {
		responses.ERROR(w, 400, err)
		return
	}
	err = m.VerifyPassword(loginDB.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		responses.ERROR(w, 400, errors.New("Invalid password"))
		return
	}
	responses.JSONString(w, 200, "Login successful")
}

func PrepareLogin(a *m.Account) {
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.Email = strings.ToLower(a.Email)
}