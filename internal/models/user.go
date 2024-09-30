package models

type User struct {
	Id           int      `json:"-" `
	Login        string   `json:"login" binding:"required"`
	Email        string   `json:"email" binding:"required"`
	HashPassword string   `json:"password" binding:"required"`
	Phone        string   `json:"phone"`
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	City         string   `json:"city"`
	Addresses    []string `json:"addresses"`
}

type UserInfo struct {
	Name      string   `json:"name"`
	Surname   string   `json:"surname"`
	Phone     string   `json:"phone"`
	City      string   `json:"city"`
	Email     string   `json:"email"`
	Addresses []string `json:"addresses"`
}
