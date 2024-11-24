package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type WorkerService struct {
	db *db.Repository
}

func (w *WorkerService) CreateWorker(login, name, surname, address, phoneNumber, password string) (int, error) {
	return w.db.CreateWorker(login, name, surname, address, phoneNumber, password)
}

func (w *WorkerService) GetWorkerId(login string) (int, error) {
	return w.db.GetWorkerId(login)
}

func (w *WorkerService) IsCorrectWorkerPassword(login string, password string) (bool, error) {
	return w.db.IsCorrectWorkerPassword(login, password)
}

func (w *WorkerService) IfWorkerExists(login string) (bool, error) {
	return w.db.IfWorkerExists(login)
}

func (w *WorkerService) ChangeWorkerCredentials(id int, login, name, surname, address, phone string) error {
	return w.db.ChangeWorkerCredentials(id, login, name, surname, address, phone)
}

func (w *WorkerService) ChangeWorkerPassword(id int, password string) error {
	return w.db.ChangeWorkerPassword(id, password)
}

func (w *WorkerService) DeleteWorker(id int) error {
	return w.db.DeleteWorker(id)
}

func (w *WorkerService) IsCorrectWorkerPasswordId(id int, passwordToCheck string) (bool, error) {
	return w.db.IsCorrectWorkerPasswordId(id, passwordToCheck)
}

func (w *WorkerService) GetWorkerInfo(id int) (models.WorkerInfo, error) {
	return w.db.GetWorkerInfo(id)
}

func (w *WorkerService) UpdateWorkerPhoto(photoString string, userId int) error {
	return w.db.UpdateWorkerPhoto(photoString, userId)
}

func NewWorkerService(db *db.Repository) *WorkerService {
	return &WorkerService{
		db: db,
	}
}
