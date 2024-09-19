package models

type Restaurant struct {
	Id          string   `json:"-"`
	Name        string   `json:"restaurant_name" binding:"required"`
	Rate        int      `json:"restaurant_rate"`
	Description string   `json:"restaurant_description" binding:"required"`
	Phone       string   `json:"restaurant_phone" binding:"required"`
	Address     string   `json:"restaurant_address" binding:"required"`
	Socials     []string `json:"restaurant_socials" binding:"required"`
	PhotoUrl    string   `json:"restaurant_photo_url" binding:"required"`
}
