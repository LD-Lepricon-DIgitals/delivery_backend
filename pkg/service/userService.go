package service

import (
	"errors"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserService struct {
	repo *db.Repository
}

func NewUserService(repo *db.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(login, name, surname, address, phoneNumber, password string) (int, error) {
	return u.repo.CreateUser(login, name, surname, address, phoneNumber, password)
}
func (u *UserService) GetUserId(login string) (int, error) {
	return u.repo.GetUserId(login)
}
func (u *UserService) IsCorrectPassword(id int, password string) (bool, error) {
	return u.repo.IsCorrectPassword(id, password)
}
func (u *UserService) IfUserExists(login string) (bool, error) { //TODO: id
	return u.repo.IfUserExists(login)
}

func (u *UserService) ChangeUserCredentials(id int, login, name, surname, address, phone string) error {

	return u.repo.ChangeUserCredentials(id, login, name, surname, address, phone)
}

func (u *UserService) ChangePassword(login string, password string) error { //TODO: id
	userId, err := u.GetUserId(login)
	if err != nil {
		return err
	}
	ok, err := u.IsCorrectPassword(userId, password)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid password")
	}
	return u.repo.ChangePassword(userId, password)

}

func (u *UserService) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
