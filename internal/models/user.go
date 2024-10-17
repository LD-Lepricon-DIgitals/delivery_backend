package models

type User struct {
	UserID   int    `db:"user_id"`
	Login    string `db:"user_login"`
	Password string `db:"user_hashed_password"`
	Phone    string `db:"user_phone"`
	Name     string `db:"user_name"`
	Surname  string `db:"user_surname"`
	Address  string `db:"user_adress"`
}

type UserInfo struct {
	UserLogin string `json:"user_login" binding:"required"`
	Phone     string `json:"user_phone" binding:"required"`
	Name      string `json:"user_name" binding:"required"`
	Surname   string `json:"user_surname" binding:"required"`
	Address   string `json:"user_adress" binding:"required"`
}

type UserReg struct {
	UserLogin   string `json:"user_login" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
	UserSurname string `json:"user_surname" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	UserPhone   string `json:"user_phone" binding:"required"`
	UserPass    string `json:"user_password" binding:"required"`
}
