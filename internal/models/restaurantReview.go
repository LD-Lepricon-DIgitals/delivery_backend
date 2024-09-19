package models

type RestaurantReview struct {
	Id           int    `json:"-"`
	CustomerId   int    `json:"customer_id" binding:"required"`
	RestaurantId int    `json:"restaurant_id" binding:"required"`
	Rating       int    `json:"rating" binding:"required"`
	Comment      string `json:"comment" binding:"required"`
	Date         string `json:"date" binding:"required"`
}
