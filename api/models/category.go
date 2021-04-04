package models

type Category struct {

	ID	int64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
	Income bool `json:"income"`

}
