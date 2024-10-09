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
	tx.Commit()

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

func (u *UserService) ChangeUserCredentials(id int, login, name, surname, address, phone string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.Exec("UPDATE users SET user_login = $1 WHERE id = $2", login, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into users table: %w", err)
	}
	_, err = tx.Exec("UPDATE users_info SET user_name = $1, user_surname = $2, user_adress = $3 WHERE user_id = $4", name, surname, address, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into users_info table: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	return nil
}

func (u *UserService) ChangePassword(id, password string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.Exec("UPDATE users SET user_hashed_password = $1 WHERE id = $2", password, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update user password: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	return nil
}
