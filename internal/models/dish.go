package models

type Dish struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"dish_name" db:"dish_name" validation:"required"`
	Description string  `json:"dish_description" db:"dish_description" validation:"required"`
	Price       float64 `json:"dish_price" db:"dish_price" validation:"required"`
	Weight      float64 `json:"dish_weight" db:"dish_weight" validation:"required"`
	Photo       string  `json:"dish_photo" db:"dish_photo" validation:"required"`
	Rating      int     `json:"dish_rating" db:"dish_rating" validation:"required"`
	Category    string  `json:"dish_category" db:"dish_category" validation:"required"`
}

type ChangeDishPayload struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"dish_name" binding:"required"`
	Price       float64 `json:"dish_price" binding:"required"`
	Weight      float64 `json:"dish_weight" binding:"required"`
	Description string  `json:"dish_description" binding:"required"`
	Rating      int     `json:"dish_rating" validation:"required"`
	Photo       string  `json:"dish_photo" binding:"required"`
	Category    string  `json:"dish_category" binding:"required"`
}
