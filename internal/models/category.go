package models

type Category struct {
	Id           int    `json:"-" db:"id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
