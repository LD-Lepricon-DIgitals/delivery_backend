package db

import (
	"errors"
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserSrv struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserSrv {
	return &UserSrv{
		db: db,
	}
}

func (u *UserSrv) Create(email, login, password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (email, login, password) VALUES ($1, $2, $3) RETURNING id;")
	row := u.db.QueryRow(query, email, login, password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("failed to create user")
	}
	return id, nil
}

func (u *UserSrv) GetId(login, password string) (int, error) {
	var id int
	query := fmt.Sprintf(`SELECT id FROM users WHERE user_login=$1 AND password=$2;`)
	row := u.db.QueryRow(query, login, password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("failed to get user id")
	}
	return id, nil
}

func (u *UserSrv) CheckIfExists(email string) (bool, error) {
	var res int
	query := fmt.Sprintf(`SELECT COUNT(1) FROM users WHERE user_email=$1;`)
	row := u.db.QueryRow(query, email)
	if err := row.Scan(&res); err != nil {
		return false, errors.New("failed to check")
	}
	return res == 1, nil
}

func (u *UserSrv) GetById(id int) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT u.id, u.user_login, u.user_email, ui.user_phone, ui.user_name, ui.user_surname, ui.user_city FROM users u LEFT JOIN users_info ui ON u.id = ui.user_id WHERE u.id = $1;")
	row := u.db.QueryRow(query, id)
	if err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Phone,
		&user.Name,
		&user.Surname,
		&user.City,
	); err != nil {
		return nil, errors.New("failed to get user info")
	}
	return &user, nil
}

func (u *UserSrv) ChangeCity(id int, city string) error {
	query := fmt.Sprintf("UPDATE users_info SET user_city=$1 WHERE user_id=$2;")
	_, err := u.db.Exec(query, city, id)
	if err != nil {
		return errors.New("failed to change user city")
	}
	return nil
}

func (u *UserSrv) ChangeLogin(id int, login string) error {
	query := fmt.Sprintf("UPDATE users SET user_login=$1 WHERE user_id=$2;")
	_, err := u.db.Exec(query, login, id)
	if err != nil {
		return errors.New("failed to change user login")
	}
	return nil
}

func (u *UserSrv) ChangePassword(id int, password string) error {
	query := fmt.Sprintf("UPDATE users SET user_password=$1 WHERE user_id=$2;")
	_, err := u.db.Exec(query, password, id)
	if err != nil {
		return errors.New("failed to change user password")
	}
	return nil
}

func (u *UserSrv) ChangeEmail(id int, email string) error {
	query := fmt.Sprintf("UPDATE users SET user_email=$1 WHERE user_id=$2;")
	_, err := u.db.Exec(query, email, id)
	if err != nil {
		return errors.New("failed to change user email")
	}
	return nil
}
