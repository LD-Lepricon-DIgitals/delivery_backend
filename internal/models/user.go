package models

type User struct {
	Login    string `db:"user_login"`
	Password string `db:"user_hashed_password"`
}

type UserInfo struct {
	UserID  int    `db:"user_id"`
	Phone   string `db:"user_phone"`
	Name    string `db:"user_name"`
	Surname string `db:"user_surname"`
	Address string `db:"user_adress"`
}
