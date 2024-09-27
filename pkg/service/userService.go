package service

import (
	"errors"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type UserService struct {
	repo *db.Repository
}

func NewUserService(repo *db.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(email, login, password string) (int, error) {
	id, err := u.repo.Create(email, login, password)
	if err != nil {
		return 0, errors.New("failed to create user")
	}

	return id, nil
}

func (u *UserService) CheckIfExists(email string) (bool, error) {
	ok, err := u.repo.CheckIfExists(email)
	if err != nil {
		return false, errors.New("failed to check user")
	}
	if !ok {
		return false, errors.New("user does not exist")
	}
	return true, nil
}

func (u *UserService) DeleteUser(id int) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

func (u *UserService) ChangeCity(id int, city string) error {
	err := u.repo.ChangeCity(id, city)
	if err != nil {
		return errors.New("failed to change city")
	}
	return nil
}

func (u *UserService) ChangeLogin(id int, login string) error {
	err := u.repo.ChangeLogin(id, login)
	if err != nil {
		return errors.New("failed to change login")
	}
	return nil
}

func (u *UserService) ChangePassword(id int, oldPassword, newPassword string) error {
	user, err := u.repo.GetById(id)
	pass, err := u.repo.GetUserPass(user.Login)

	if err != nil {
		return errors.New("failed to find user")
	}

	if oldPassword != pass {
		return errors.New("invalid password")
	}
	err = u.repo.ChangePassword(id, newPassword)
	if err != nil {
		return errors.New("failed to change password")
	}
	return nil
}

func (u *UserService) ChangeEmail(id int, email string) error {
	err := u.repo.ChangeEmail(id, email)
	if err != nil {
		return errors.New("failed to change email")
	}
	return nil
}
func (u *UserService) GetById(id int) (*models.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return nil, errors.New("failed to get user")
	}
	return user, nil
}

func (u *UserService) ChangePhone(id int, phone string) error {
	err := u.repo.ChangePhone(id, phone)
	if err != nil {
		return errors.New("failed to change user phone number")
	}
	return nil
}

func (u *UserService) AddUserInfo(id int, userPhone, userName, userSurname, userCity string) error {
	err := u.repo.AddUserInfo(id, userPhone, userName, userSurname, userCity)
	if err != nil {
		return errors.New("failed to add user info")
	}
	return nil
}

func (u *UserService) AddUserAddress(id int, address string) error {
	err := u.repo.AddUserAddress(id, address)
	if err != nil {
		return errors.New("failed to add user address")
	}
	return nil
}
