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
	UserID  int    `db:"user_id"`
	Phone   string `db:"user_phone"`
	Name    string `db:"user_name"`
	Surname string `db:"user_surname"`
	Address string `db:"user_adress"`
}

type UserReg struct {
	UserLogin   string `json:"user_login"`
	UserName    string `json:"user_name"`
	UserSurname string `json:"user_surname"`
	UserAddress string `json:"user_address"`
	UserPhone   string `json:"user_phone"`
	UserPass    string `json:"user_password"`
}
