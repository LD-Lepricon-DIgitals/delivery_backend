package models

import "database/sql"

type Dish struct {
	Id          int             `json:"-"`
	Name        string          `json:"dish_name" binding:"required"`
	Description string          `json:"dish_description" binding:"required"`
	Price       float64         `json:"dish_price" binding:"required"`
	Weight      float64         `json:"dish_weight" binding:"required"`
	PhotoUrl    string          `json:"dish_photo_url" binding:"required"`
	Rating      sql.NullFloat64 `json:"dish_rating"`
	Category    string          `json:"dish_category"`
}
