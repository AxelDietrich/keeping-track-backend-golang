package models

import (
	"html"
	"strings"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Account) BeforeSave() error {
	a.Username = html.EscapeString(strings.TrimSpace(a.Username))
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.Email = strings.ToLower(a.Email)
	a.CreatedAt = time.Now()
	hashedPassword, err := Hash(a.Password)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}
