package repositories

import (
	"gorm.io/gorm"
	"keeping-track-backend-golang/api/models"
)

func CreateAccount(db *gorm.DB, account *models.Account){
	db.Create(account)
}