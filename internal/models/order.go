package models

type Order struct {
	Id         int  `json:"-"`
	CustomerId int  `json:"customer_id" validation:"required"`
	WorkerId   *int `json:"worker_id"`
	//RestaurantId int         `json:"restaurant_id" validation:"required"`
	Price  float64     `json:"order_price" validation:"required"`
	Status string      `json:"order_status"`
	Dishes []OrderDish `json:"dishes" validation:"required"` // Используем слайс OrderDish
}

type CreateOrder struct {
	CustomerId int         `json:"customer_id" validation:"required"`
	Dishes     []OrderDish `json:"dishes" validation:"required"`
	Price      float64     `json:"order_price" validation:"required"`
}

type OrderDetails struct {
	UserName     string      `json:"username" validation:"required"`
	UserSurname  string      `json:"user_surname" validation:"required"`
	UserLogin    string      `json:"user_login" validation:"required"`
	UserPhotoUrl string      `json:"user_photo_url" validation:"required"`
	OrderPrice   string      `json:"order_price" validation:"required"`
	OrderStatus  string      `json:"order_status" validation:"required"`
	Dishes       []OrderDish `json:"dishes" validation:"required"`
}

type OrderInfo struct {
	OrderId     int    `json:"order_id" validation:"required"`
	UserLogin   string `json:"user_login" validation:"required"`
	UserPhoto   string `json:"user_photo" validation:"required"`
	OrderStatus string `json:"order_status" validation:"required"`
}
