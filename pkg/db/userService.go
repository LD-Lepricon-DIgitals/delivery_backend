package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(login, name, surname, address, phoneNumber, password string) (int, error) {
	var userId int

	tx, err := u.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	err = tx.QueryRow("INSERT INTO users (user_login, user_hashed_password) VALUES ($1, $2) RETURNING id;", login, password).Scan(&userId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into users table: %w", err)
	}
	_, err = u.db.Exec("INSERT INTO users_info (user_id, user_phone, user_name, user_surname, user_adress) VALUES ($1, $2, $3, $4, $5)", userId, phoneNumber, name, surname, address)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into users_info table: %w", err)
	}

	log.Println(fmt.Sprintf("user %d created", userId))
	return userId, nil
}

func (u *UserService) GetUserId(login string) (int, error) {
	var userId int
	err := u.db.Get(&userId, "SELECT id FROM users WHERE user_login = $1", login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get user id: %w", err)
	}
	return userId, nil
}

func (u *UserService) IfUserExists(login string) (bool, error) {
	var userId int
	err := u.db.Get(&userId, "SELECT id FROM users WHERE user_login = $1", login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("error checking user existance: %w", err)
	}
	return true, nil
}

func (u *UserService) IsCorrectPassword(login, password string) (bool, error) {
	var userId int
	err := u.db.Get(&userId, "SELECT id FROM users WHERE user_login = $1 AND user_hashed_password=$2", login, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("error checking password: %w", err)
	}
	return true, nil
}
