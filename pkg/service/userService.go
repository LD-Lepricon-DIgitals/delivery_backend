package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserService struct {
	repo *db.Repository
}

func NewUserService(repo *db.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(login, name, surname, address, phoneNumber, password string) (int, error) {
	userId, err := u.repo.CreateUser(login, name, surname, address, phoneNumber, password)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
func (u *UserService) GetUserId(login string) (int, error) {
	userId, err := u.repo.GetUserId(login)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
func (u *UserService) IsCorrectPassword(login, password string) (bool, error) {
	correctPassword, err := u.repo.IsCorrectPassword(login, password)
	if err != nil {
		return false, err
	}
	return correctPassword, nil
}
func (u *UserService) IfUserExists(username string) (bool, error) {
	exists, err := u.repo.IfUserExists(username)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserService) ChangeUserCredentials(id int, login, name, surname, address, phone string) error {
	err := u.repo.ChangeUserCredentials(id, login, name, surname, address, phone)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ChangePassword(id int, password string) error {
	user
	ok, err := u.IsCorrectPassword(id, password)
	err := u.repo.ChangePassword(id, password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(id int) error {
	return nil
}
