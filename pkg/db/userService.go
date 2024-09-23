package db

import "github.com/jmoiron/sqlx"

type UserSrv struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserSrv {
	return &UserSrv{
		db: db,
	}
}

func (u *UserSrv) GetId(username, password string) (int, error) {
	return 0, nil
}
