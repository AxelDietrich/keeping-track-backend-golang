package models

import (
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type Account struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (Account) TableName() string {
	return "keepingtrack.accounts"
}

func (a *Account) BeforeSave(db *gorm.DB) error {
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
