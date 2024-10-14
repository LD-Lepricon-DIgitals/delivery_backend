package models

type Dish struct {
	Id          int     `json:"-"`
	Name        string  `json:"dish_name" binding:"required"`
	Description string  `json:"dish_description" binding:"required"`
	Price       float64 `json:"dish_price" binding:"required"`
	Weight      float64 `json:"dish_weight" binding:"required"`
	PhotoUrl    string  `json:"dish_photo_url" binding:"required"`
	Rating      float64 `json:"dish_rating"`
	Category    string  `json:"dish_category"`
}
