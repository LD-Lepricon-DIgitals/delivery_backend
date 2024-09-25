package service

import (
	"errors"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserService struct {
	repo db.Repository
}

func NewUserService(repo db.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(email, login, password string) (int, error) {
	id, err := u.repo.Create(email, login, password)
	if err != nil {
		return 0, errors.New("Failed to create user")
	}

	return id, nil
}

func (u *UserService) CheckIfExists(email string) (bool, error) {
	ok, err := u.repo.CheckIfExists(email)
	if err != nil {
		return false, errors.New("Failed to check user")
	}
	if !ok {
		return false, errors.New("User does not exist")
	}
	return true, nil
}

func (u *UserService) DeleteUser(email string) error {
	return nil //TODO: implement
}

func (u *UserService) ChangeCity(id int, city string) error {
	err := u.repo.ChangeCity(id, city)
	if err != nil {
		return errors.New("Failed to change city")
	}
	return nil
}

func (u *UserService) ChangeLogin(id int, login string) error {
	err := u.repo.ChangeLogin(id, login)
	if err != nil {
		return errors.New("Failed to change login")
	}
	return nil
}

func (u *UserService) ChangePassword(id int, password string) error {
	err := u.repo.ChangePassword(id, password)
	if err != nil {
		return errors.New("Failed to change password")
	}
	return nil
}

func (u *UserService) ChangeEmail(id int, email string) error {
	err := u.repo.ChangeEmail(id, email)
	if err != nil {
		return errors.New("Failed to change email")
	}
	return nil
}
