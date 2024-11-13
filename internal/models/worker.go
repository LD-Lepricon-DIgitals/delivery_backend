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

type WorkerInfo struct {
	WorkerLogin string `json:"worker_login" validate:"required"`
	Phone       string `json:"worker_phone" validate:"required"`
	Name        string `json:"worker_name" validate:"required"`
	Surname     string `json:"worker_surname" validate:"required"`
	Address     string `json:"worker_address" validate:"required"`
	Photo       string `json:"worker_photo"`
}
