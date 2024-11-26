package models

type User struct {
	UserID   int    `db:"user_id"`
	Login    string `db:"user_login"`
	Password string `db:"user_hashed_password"`
	Phone    string `db:"user_phone"`
	Name     string `db:"user_name"`
	Surname  string `db:"user_surname"`
	Address  string `db:"user_adress"`
	Photo    string `db:"user_photo"`
	Role     string `db:"user_role"`
}

type UserInfo struct {
	UserLogin string `json:"user_login" validate:"required"`
	Phone     string `json:"user_phone" validate:"required"`
	Name      string `json:"user_name" validate:"required"`
	Surname   string `json:"user_surname" validate:"required"`
	Address   string `json:"user_address" validate:"required"`
	Photo     string `json:"user_photo" validate:"required"`
	Role      string `json:"user_role" validate:"required"`
}

type UserReg struct {
	UserLogin   string `json:"user_login" validate:"required"`
	UserName    string `json:"user_name" validate:"required"`
	UserSurname string `json:"user_surname" validate:"required"`
	UserAddress string `json:"user_address" validate:"required"`
	UserPhone   string `json:"user_phone" validate:"required"`
	UserPass    string `json:"user_password" validate:"required"`
	Role        string `json:"user_role" validate:"required"`
}

type ChangeUserCredsPayload struct {
	UserLogin string `json:"user_login" validate:"required"`
	Phone     string `json:"user_phone" validate:"required"`
	Name      string `json:"user_name" validate:"required"`
	Surname   string `json:"user_surname" validate:"required"`
	Address   string `json:"user_address" validate:"required"`
}
