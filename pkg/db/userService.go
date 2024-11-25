package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(user models.UserReg) (int, error) {
	var userId int

	tx, err := u.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Вставка пользователя
	err = tx.QueryRow("INSERT INTO users (user_login, user_hashed_password,user_role) VALUES ($1, $2) RETURNING id;", user.UserLogin, user.UserPass, user.Role).Scan(&userId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into users table: %w", err)
	}

	// Вставка информации о пользователе
	_, err = tx.Exec("INSERT INTO users_info (user_id, user_phone, user_name, user_surname, user_address,user_photo) VALUES ($1, $2, $3, $4, $5, $6)", userId, user.UserPhone, user.UserName, user.UserSurname, user.UserAddress, "") //TODO: Add default photo
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into users_info table: %w", err)
	}

	// Завершение транзакции
	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
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

func (u *UserService) IsCorrectPassword(login string, passwordToCheck string) (bool, error) {
	var password string
	err := u.db.Get(&password, "SELECT user_hashed_password FROM users WHERE user_login = $1", login)
	if err != nil {
		return false, fmt.Errorf("error checking password: %w", err)
	}
	if password != passwordToCheck {
		return false, nil
	}
	return true, nil
}

func (u *UserService) ChangeUserCredentials(id int, info models.UserInfo) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	//Updating users
	res, err := tx.Exec("UPDATE users SET user_login = $1 WHERE id = $2", info.UserLogin, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into users table: %w", err)
	}
	affect, _ := res.RowsAffected()
	if affect == 0 {
		tx.Rollback()
		return fmt.Errorf("user %d not found", id)
	}

	//Updating users_info
	res, err = tx.Exec("UPDATE users_info SET user_name = $1, user_surname = $2, user_address = $3, user_phone = $4 WHERE user_id = $4", info.Name, info.Surname, info.Address, info.Phone, id)
	if err != nil {

		tx.Rollback()
		return fmt.Errorf("failed to insert into users_info table: %w", err)
	}
	//Affected rows checking
	affect, _ = res.RowsAffected()
	if affect == 0 {
		tx.Rollback()
		return fmt.Errorf("user %d not found", id)
	}
	//Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	return nil
}

func (u *UserService) ChangePassword(id int, password string) error {
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

func (u *UserService) DeleteUser(id int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete user: %w", err)
	}
	_, err = tx.Exec("DELETE FROM users_info WHERE user_id = $1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete user: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit user deletion: %w", err)
	}
	return nil
}

func (u *UserService) IsCorrectPasswordId(id int, passwordToCheck string) (bool, error) {
	var password string
	err := u.db.Get(&password, "SELECT user_hashed_password FROM users WHERE id = $1", id)
	if err != nil {
		return false, fmt.Errorf("error checking password: %w", err)
	}
	if password != passwordToCheck {
		return false, nil
	}
	return true, nil
}

func (u *UserService) GetUserInfo(id int) (models.UserInfo, error) {
	var user models.UserInfo
	tx, err := u.db.Begin()
	if err != nil {
		return user, fmt.Errorf("failed to begin transaction: %w", err)
	}
	err = tx.QueryRow("SELECT user_login FROM users WHERE id = $1", id).Scan(&user.UserLogin)
	if err != nil {
		return user, fmt.Errorf("failed to get user info: %w", err)
	}
	err = tx.QueryRow("SELECT user_phone, user_name, user_surname, user_address, user_photo FROM users_info WHERE user_id = $1", id).Scan(&user.Phone, &user.Name, &user.Surname, &user.Address, &user.Photo)
	if err != nil {
		return user, fmt.Errorf("failed to get user info: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return user, fmt.Errorf("failed to get user info: %w", err)
	}
	return user, nil
}

func (u *UserService) UpdatePhoto(photoString string, userId int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("UPDATE users_info SET user_photo=$1 WHERE user_id = $2", photoString, userId)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return errors.New("failed to update photo")
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
