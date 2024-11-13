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
		return 0, fmt.Errorf("failed to get worker id: %w", err)
	}
	return workerId, nil
}

func (w *WorkerService) IsCorrectWorkerPassword(login string, passwordCheck string) (bool, error) {
	var password string
	err := w.db.Get(&password, "SELECT worker_hashed_password FROM workers WHERE worker_login = $1", login)
	if err != nil {
		return false, fmt.Errorf("error checking password: %w", err)
	}
	if password != passwordCheck {
		return false, nil
	}
	return true, nil
}

func (w *WorkerService) IfWorkerExists(login string) (bool, error) {
	var workerId int
	err := w.db.Get(&workerId, "SELECT id FROM workers WHERE worker_login = $1", login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("error checking worker existance: %w", err)
	}
	return true, nil
}

func (w *WorkerService) ChangeWorkerCredentials(id int, login, name, surname, address, phone string) error {
	tx, err := w.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	//Updating workers
	res, err := tx.Exec("UPDATE workers SET worker_login = $1 WHERE id = $2", login, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into workers table: %w", err)
	}
	affect, _ := res.RowsAffected()
	if affect == 0 {
		tx.Rollback()
		return fmt.Errorf("worker %d not found", id)
	}

	//Updating workers_info
	res, err = tx.Exec("UPDATE workers_info SET worker_name = $1, worker_surname = $2, worker_address = $3 WHERE worker_id = $4", name, surname, address, id)
	if err != nil {

		tx.Rollback()
		return fmt.Errorf("failed to insert into workers_info table: %w", err)
	}
	//Affected rows checking
	affect, _ = res.RowsAffected()
	if affect == 0 {
		tx.Rollback()
		return fmt.Errorf("worker %d not found", id)
	}
	//Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	return nil
}

func (w *WorkerService) ChangeWorkerPassword(id int, password string) error {
	tx, err := w.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.Exec("UPDATE workers SET worker_hashed_password = $1 WHERE id = $2", password, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update worker password: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	return nil
}

func (w *WorkerService) DeleteWorker(id int) error {
	tx, err := w.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.Exec("DELETE FROM workers WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete worker: %w", err)
	}
	_, err = tx.Exec("DELETE FROM workers_info WHERE worker_id = $1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete worker: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit worker deletion: %w", err)
	}
	return nil
}

func (w *WorkerService) IsCorrectWorkerPasswordId(id int, passwordToCheck string) (bool, error) {
	var password string
	err := w.db.Get(&password, "SELECT worker_hashed_password FROM workers WHERE id = $1", id)
	if err != nil {
		return false, fmt.Errorf("error checking password: %w", err)
	}
	if password != passwordToCheck {
		return false, nil
	}
	return true, nil
}

func (w *WorkerService) GetWorkerInfo(id int) (models.WorkerInfo, error) {
	var worker models.WorkerInfo
	tx, err := w.db.Begin()
	if err != nil {
		return worker, fmt.Errorf("failed to begin transaction: %w", err)
	}
	err = tx.QueryRow("SELECT worker_login FROM workers WHERE id = $1", id).Scan(&worker.WorkerLogin)
	if err != nil {
		return worker, fmt.Errorf("failed to get worker info: %w", err)
	}
	err = tx.QueryRow("SELECT worker_phone, worker_name, worker_surname, worker_address, worker_photo FROM workers_info WHERE worker_id = $1", id).Scan(&worker.Phone, &worker.Name, &worker.Surname, &worker.Address, &worker.Photo)
	if err != nil {
		return worker, fmt.Errorf("failed to get worker info: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return worker, fmt.Errorf("failed to get worker info: %w", err)
	}
	return worker, nil
}

func (w *WorkerService) UpdateWorkerPhoto(photoString string, workerId int) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("UPDATE workers_info SET worker_photo=$1 WHERE worker_id = $2", photoString, workerId)
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
	_, err = tx.Exec("INSERT INTO workers_info (worker_id, worker_phone, worker_name, worker_surname, worker_address,worker_photo) VALUES ($1, $2, $3, $4, $5,&6)", workerId, phoneNumber, name, surname, address, "")
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert into workers_info table: %w", err)
	}

	// Завершение транзакции
	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println(fmt.Sprintf("worker %d created", workerId))
	return workerId, nil
}
