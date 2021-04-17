package controllers

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"html"
	"io/ioutil"
	m "keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/repositories"
	"keeping-track-backend-golang/api/requests"
	"keeping-track-backend-golang/api/responses"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("gopher4life")

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	login := &requests.Login{}
	err = json.Unmarshal(reqBody, login)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	PrepareLogin(login)
	loginDB, err := repositories.GetAccount(server.DB, login)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = m.VerifyPassword(loginDB.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Invalid password"))
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: login.Email,
		UserID:   loginDB.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	responses.JSON(w, http.StatusOK, "Login successful")
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				responses.ERROR(w, http.StatusUnauthorized, err)
				return
			}
			responses.ERROR(w, http.StatusBadRequest, err)
		}
		tokenString := c.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				responses.ERROR(w, http.StatusUnauthorized, err)
				return
			}
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		if !tkn.Valid {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		r.Header.Set("userID", string(claims.UserID))

		next.ServeHTTP(w, r)
	})

}
func PrepareLogin(a *requests.Login) {
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.Email = strings.ToLower(a.Email)
}
