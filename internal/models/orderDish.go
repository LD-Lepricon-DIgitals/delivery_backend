package models

type OrderDish struct {
	DishId   int    `json:"dish_id" validation:"required"`
	DishName string `json:"dish_name"`
	Quantity int    `json:"quantity" validation:"required"`
}
