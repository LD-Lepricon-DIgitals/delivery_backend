package models

type Order struct {
	Id           int         `json:"-"`
	CustomerId   int         `json:"customer_id" binding:"required"`
	WorkerId     int         `json:"worker_id"`
	RestaurantId int         `json:"restaurant_id" binding:"required"`
	Price        float64     `json:"order_price" binding:"required"`
	Dishes       []OrderDish `json:"dishes" binding:"required"` // Используем слайс OrderDish
}

type CreateOrder struct {
	CustomerId int         `json:"customer_id" validation:"required"`
	Dishes     []OrderDish `json:"dishes" validation:"required"`
	Price      float64     `json:"order_price" validation:"required"`
}
