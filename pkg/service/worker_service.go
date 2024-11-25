package service

import "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"

type WorkerService struct {
	db *db.Repository
}

func (w WorkerService) ConfirmOrder(orderId int, workerId int) error {
	return w.ConfirmOrder(orderId, workerId)
}

func (w WorkerService) DeclineOrder(orderId int) error {
	return w.DeclineOrder(orderId)
}

func NewWorkerService(db *db.Repository) *WorkerService {
	return &WorkerService{}
}
