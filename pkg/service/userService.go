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
func (u *UserService) GetUserId(username, password string) (int, error) {
	return 0, nil
}
func (u *UserService) IsCorrectPassword(login, password string) (bool, error) {
	return false, nil
}
func (u *UserService) IfUserExists(username string) (bool, error) {
	return false, nil
}
