package models

type Worker struct {
	Id           int    `json:"-" `
	Login        string `json:"login" binding:"required"`
	Email        string `json:"email" binding:"required"`
	HashPassword string `json:"password" binding:"required"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	City         string `json:"city"`
}
