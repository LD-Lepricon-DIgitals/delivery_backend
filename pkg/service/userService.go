package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
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
func (u *UserService) IsCorrectPassword(login string, password string) (bool, error) {
	return u.repo.IsCorrectPassword(login, password)
}
func (u *UserService) IfUserExists(login string) (bool, error) { //TODO: id
	return u.repo.IfUserExists(login)
}

func (u *UserService) ChangeUserCredentials(id int, login, name, surname, address, phone string) error {

	return u.repo.ChangeUserCredentials(id, login, name, surname, address, phone)
}

func (u *UserService) ChangePassword(id int, password string) error {
	return u.repo.ChangePassword(id, password)

}

func (u *UserService) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}

func (u *UserService) IsCorrectPasswordId(id int, passwordToCheck string) (bool, error) {
	return u.repo.IsCorrectPasswordId(id, passwordToCheck)
}

func (u *UserService) GetUserInfo(id int) (models.UserInfo, error) {
	return u.repo.GetUserInfo(id)
}

func (u *UserService) UpdatePhoto(photo string, userId int) error {
	return u.repo.UpdatePhoto(photo, userId)
}
