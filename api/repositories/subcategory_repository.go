package repositories

import (
	"errors"
	"gorm.io/gorm"
	"keeping-track-backend-golang/api/models"
)

func CreateSubcategory(db *gorm.DB, s *models.Subcategory) error {

	var err error
	var sub *models.Subcategory
	err = db.Where("name = ? AND category_id = ?", s.Name, s.CategoryID).First(&sub).Error
	if err == nil {
		return errors.New("There is already a subcategory with that name")
	}
	err = db.Create(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateSubcategory(db *gorm.DB, id int, name string) error {

	var err error
	s := &models.Subcategory{}
	s.ID = id
	err = db.First(&s).Error
	if err != nil {
		return err
	}
	s.Name = name
	err = db.Save(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSubcategory(db *gorm.DB, subID int) (*models.Subcategory, error) {

	var err error
	var sub *models.Subcategory
	err = db.Find(&sub).Where("id = ?", subID).Error
	if err != nil {
		return &models.Subcategory{}, err
	}
	return sub, err
}

func GetAllSubcategories(db *gorm.DB, categoryID int) ([]*models.Subcategory, error) {
	var err error
	var subcategories []*models.Subcategory
	err = db.Find(&subcategories).Where("category_id = ?", categoryID).Error
	if err != nil {
		panic(err)
	}
	return subcategories, nil
}

func DeleteSubcategory(db *gorm.DB, subcategoryID int) error {
	var err error
	sub := &models.Subcategory{}
	err = db.Where("id = ?", subcategoryID).Delete(&sub).Error
	if err != nil {
		return err
	}
	return nil
}
