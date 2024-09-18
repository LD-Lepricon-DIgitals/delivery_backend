package models

type OrderDish struct {
	DishId   int `json:"dish_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}
