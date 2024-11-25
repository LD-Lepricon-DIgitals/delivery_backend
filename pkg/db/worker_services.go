package db

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type WorkerService struct {
	db *sqlx.DB
}

func (w WorkerService) ConfirmOrder(orderId int, workerId int) error {
	tx, err := w.db.Beginx()
	if err != nil {
		return errors.New("failed to confirm order")
	}

	rows, err := tx.Query("UPDATE orders SET worker_id = $1 WHERE id = $2 ", workerId, orderId)
	if err != nil || rows == nil {
		tx.Rollback()
		return errors.New("failed to confirm order")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("failed to confirm order")
	}
	return nil
}

func (w WorkerService) DeclineOrder(orderId int) error {
	tx, err := w.db.Beginx()
	if err != nil {
		return errors.New("failed to decline order")
	}

	rows, err := tx.Query("UPDATE orders SET worker_id = 0 WHERE id = $2 ", orderId)
	if err != nil || rows == nil {
		tx.Rollback()
		return errors.New("failed to decline order")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("failed to decline order")
	}
	return nil
}

func NewWorkerService(db *sqlx.DB) *WorkerService {
	return &WorkerService{db: db}
}
