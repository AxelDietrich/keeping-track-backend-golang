package repositories

import (
	"database/sql"
	"errors"
	"keeping-track-backend-golang/api/models"
)

func CheckIfUserSubcategory(db *sql.DB, subID int, accID int) error {
	var id int
	err := db.QueryRow("select id from keepingtrack.subcategories s inner join keepingtrack.categories c on "+
		"s.category_id = c.id where c.account_id = $1 and s.id = $2", accID, subID).
		Scan(&id)
	if err != nil || id == 0 {
		return err
	}
	return nil
}

func CreateSubcategory(db *sql.DB, s *models.Subcategory) error {

	var err error
	var isNameUsed int
	err = db.QueryRow("select count(*) from keepingtrack.subcategories where name = $1 AND category_id = $2;", s.Name, s.CategoryID).Scan(&isNameUsed)
	if err != nil {
		return err
	}
	if isNameUsed > 0 {
		return errors.New("There is already a subcategory with that name")
	}
	err = db.QueryRow("insert into keepingtrack.subcategories (name, amount, category_id) values ($1, $2, $3);", s.Name, s.Amount, s.CategoryID).Err()
	if err != nil {
		return err
	}
	return nil
}

func UpdateSubcategory(db *sql.DB, id int, name string) error {

	var err error
	sub := &models.Subcategory{}
	err = db.QueryRow("select * from keepingtrack.subcategories where id = $1;", id).Scan(&sub.ID, &sub.Name, &sub.Amount, &sub.CategoryID, &sub.CreatedAt, &sub.UpdatedAt)
	if err != nil {
		return err
	}
	err = db.QueryRow("update keepingtrack.subcategories set name = $1 where id = $2;", name, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetSubcategory(db *sql.DB, subID int) (*models.Subcategory, error) {

	var err error
	var sub *models.Subcategory
	err = db.QueryRow("select * from keepingtrack.subcategories where id = $1;", subID).Scan(&sub.ID, &sub.Name, &sub.Amount, &sub.CategoryID, &sub.CreatedAt, &sub.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return sub, err
}

func GetAllSubcategories(db *sql.DB, categoryID int) ([]*models.Subcategory, error) {
	var (
		err           error
		subcategories []*models.Subcategory
		sub           *models.Subcategory
	)

	rows, err := db.Query("select * from keepingtrack.subcategories where category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&sub.ID, &sub.Name, &sub.Amount, &sub.CategoryID, &sub.CreatedAt, &sub.UpdatedAt)
		subcategories = append(subcategories, sub)
	}

	rows.Close()
	return subcategories, nil
}

func DeleteSubcategory(db *sql.DB, subcategoryID int) error {
	var err error
	err = db.QueryRow("delete from keepingtrack.subcategories where id = $1;", subcategoryID).Err()
	if err != nil {
		return err
	}
	return nil
}
