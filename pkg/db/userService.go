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
	query := fmt.Sprintf("INSERT INTO users (user_email, user_login, user_hashed_password) VALUES ($1, $2, $3) RETURNING id;")
	row := u.db.QueryRow(query, email, login, password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("failed to create user")
	}
	return id, nil
}

func (u *UserSrv) GetId(login, password string) (int, error) {
	var id int
	query := fmt.Sprintf(`SELECT id FROM users WHERE user_login=$1 AND user_hashed_password=$2;`)
	row := u.db.QueryRow(query, login, password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("failed to get user id")
	}
	return id, nil
}

func (u *UserSrv) CheckIfExists(login string) (bool, error) {
	var res int
	query := fmt.Sprintf(`SELECT COUNT(1) FROM users WHERE user_login=$1;`)
	row := u.db.QueryRow(query, login)
	if err := row.Scan(&res); err != nil {
		return false, errors.New("failed to check")
	}
	return res == 1, nil
}

func (u *UserSrv) GetUserInfo(id int) (*models.UserInfo, error) {
	var res1, res2 int
	var user models.UserInfo
	tx, err := u.db.Begin()
	if err != nil {
		return nil, errors.New("failed to start transaction")
	}
	defer tx.Rollback()
	userInfoCountQuery := fmt.Sprintf("SELECT COUNT(1) FROM users_info WHERE user_id=$1;")
	row := tx.QueryRow(userInfoCountQuery, id)
	if err := row.Scan(&res1); err != nil {
		return nil, errors.New("failed to check userInfo")
	}
	if res1 <= 0 {
		return nil, errors.New("userInfo is empty")
	}
	userInfoQuery := fmt.Sprintf("SELECT u.user_name, u.user_surname, u.user_phone, u.user_city, us.user_email FROM users_info AS u INNER JOIN users AS us ON us.id=u.user_id WHERE u.user_id=$1;")
	row = tx.QueryRow(userInfoQuery, id)
	if err := row.Scan(&user.Name, &user.Surname, &user.Phone, &user.City, &user.Email); err != nil {
		return nil, errors.New("failed to get userInfo")
	}
	userAddressesCountQuery := fmt.Sprintf("SELECT COUNT(1) FROM user_addresses WHERE user_id=$1;")
	row = tx.QueryRow(userAddressesCountQuery, id)
	if err := row.Scan(&res2); err != nil {
		return nil, errors.New("failed to check userAddresses")
	}
	if res2 <= 0 {
		return nil, errors.New("userAddresses is empty")
	}
	userAddressesQuery := fmt.Sprintf("SELECT user_address FROM user_addresses WHERE user_id=$1;")
	rows, err := tx.Query(userAddressesQuery, id)
	if err != nil {
		return nil, errors.New("failed to get userAddresses")
	}
	defer rows.Close()

	for rows.Next() {
		var address string
		if err := rows.Scan(&address); err != nil {
			return nil, errors.New("failed to scan user address")
		}
		user.Addresses = append(user.Addresses, address)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("error iterating over user addresses")
	}
	if err := tx.Commit(); err != nil {
		return nil, errors.New("failed to commit transaction")
	}
	return &user, nil
}

func (u *UserSrv) ChangePassword(id int, password string) error {
	query := fmt.Sprintf("UPDATE users SET user_hashed_password=$1 WHERE id=$2;")
	_, err := u.db.Exec(query, password, id)
	if err != nil {
		return errors.New("failed to change user password")
	}
	return nil
}

func (u *UserSrv) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id=$1;")
	if _, err := u.db.Exec(query, id); err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

func (u *UserSrv) GetUserPass(username string) (string, error) {
	var password string
	query := fmt.Sprintf("SELECT u.user_hashed_password FROM users AS u WHERE u.user_login=$1;")
	row := u.db.QueryRow(query, username)

	if err := row.Scan(&password); err != nil {
		return "", errors.New("failed to get user password")
	}

	return password, nil
}

func (u *UserSrv) AddUserAddress(id int, address string) error {
	query := fmt.Sprintf("INSERT INTO user_addresses (user_id, user_address) VALUES ($1, $2);")
	_, err := u.db.Exec(query, id, address)
	if err != nil {
		return errors.New("failed to add user address")
	}
	return nil
}

func (u *UserSrv) ChangeUserCredentials() error {
	//TODO implement me
	panic("implement me")
}

/*func (u *UserSrv) ChangeCity(id int, city string) error {
	query := fmt.Sprintf("UPDATE users_info SET user_city = $1 WHERE user_id = $2;")
	_, err := u.db.Exec(query, city, id)
	if err != nil {
		return errors.New("failed to change user city")
	}
	return nil
}*/

/*func (u *UserSrv) ChangeLogin(id int, login string) error {
	query := fmt.Sprintf("UPDATE users SET user_login=$1 WHERE id=$2;")
	_, err := u.db.Exec(query, login, id)
	if err != nil {
		return errors.New("failed to change user login")
	}
	return nil
}
*/

/*func (u *UserSrv) ChangeEmail(id int, email string) error {
	query := fmt.Sprintf("UPDATE users SET user_email=$1 WHERE id=$2;")
	_, err := u.db.Exec(query, email, id)
	if err != nil {
		return errors.New("failed to change user email")
	}
	return nil
}
*/
/*func (u *UserSrv) ChangePhone(id int, phone string) error {
	query := fmt.Sprintf("UPDATE users_info SET user_phone=$1 WHERE user_id=$2;")
	_, err := u.db.Exec(query, phone, id)
	if err != nil {
		return errors.New("failed to change user phone number")
	}
	return nil
}*/
