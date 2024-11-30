package models

type Category struct {
	Id           int    `json:"-"`
	CategoryName string `json:"category_name"`
}
