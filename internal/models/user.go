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
