package models

type Admin struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}
