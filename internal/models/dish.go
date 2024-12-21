package models

type Dish struct {
	Id          int     `json:"id"`
	Name        string  `json:"dish_name" validation:"required"`
	Description string  `json:"dish_description" validation:"required"`
	Price       float64 `json:"dish_price" validation:"required"`
	Weight      float64 `json:"dish_weight" validation:"required"`
	Photo       string  `json:"dish_photo" validation:"required"`
	Rating      int     `json:"dish_rating" validation:"required"`
	Category    string  `json:"dish_category" validation:"required"`
}

type ChangeDishPayload struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"dish_name" binding:"required"`
	Price       float64 `json:"dish_price" binding:"required"`
	Weight      float64 `json:"dish_weight" binding:"required"`
	Description string  `json:"dish_description" binding:"required"`
	Photo       string  `json:"dish_photo" binding:"required"`
	Category    int     `json:"dish_category" binding:"required"`
}
