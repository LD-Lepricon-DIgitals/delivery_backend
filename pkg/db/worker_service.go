package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type WorkerService struct {
	db *sqlx.DB
}

func (w *WorkerService) GetWorkerId(login string) (int, error) {
	var workerId int
	err := w.db.Get(&workerId, "SELECT id FROM workers WHERE worker_login = $1", login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get user id: %w", err)
	}
	return workerId, nil
}

func (w *WorkerService) IsCorrectWorkerPassword(login string, password string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) IfWorkerExists(login string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) ChangeWorkerCredentials(id int, login, name, surname, address, phone string) error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) ChangeWorkerPassword(id int, password string) error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) DeleteWorker(id int) error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) IsCorrectWorkerPasswordId(id int, passwordToCheck string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) GetWorkerInfo(id int) (models.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerService) UpdateWorkerPhoto(photoString string, userId int) error {
	//TODO implement me
	panic("implement me")
}

func NewWorkerService(db *sqlx.DB) *WorkerService {
	return &WorkerService{db: db}
}

func (w *WorkerService) CreateWorker(login, name, surname, address, phoneNumber, password string) (int, error) {
	var workerId int

	tx, err := w.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Вставка пользователя
	err = tx.QueryRow("INSERT INTO workers (worker_login, worker_hashed_password) VALUES ($1, $2) RETURNING id;", login, password).Scan(&workerId)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into workers table: %w", err)
	}

	// Вставка информации о пользователе
	_, err = tx.Exec("INSERT INTO users_info (user_id, user_phone, user_name, user_surname, user_address,user_photo) VALUES ($1, $2, $3, $4, $5,&6)", workerId, phoneNumber, name, surname, address, "")
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into users_info table: %w", err)
	}

	// Завершение транзакции
	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println(fmt.Sprintf("user %d created", workerId))
	return workerId, nil
}
