package models

import (
	"time"
)

type Account struct {

	Id	int `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

}

func (Account) TableName() string {
	return "keepingtrack.accounts"
}
