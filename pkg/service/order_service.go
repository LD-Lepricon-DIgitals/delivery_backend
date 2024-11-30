package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type OrderService struct {
	repo *db.Repository
}

func (o OrderService) GetWorkerOrders(workerId int) ([]models.Order, error) {
	return o.repo.GetWorkerOrders(workerId)
}

func (o OrderService) CreateOrder(order models.CreateOrder) (int, error) {
	return o.repo.CreateOrder(order)
}

func (o OrderService) GetOrder(orderId int) (models.Order, error) {
	return o.repo.GetOrder(orderId)
}

func (o OrderService) DeleteOrder(orderId int) error {
	return o.repo.DeleteOrder(orderId)
}

func (o OrderService) GetUsersOrders(i int) ([]models.Order, error) {
	return o.repo.GetUsersOrders(i)
}

func (o OrderService) GetOrderCustomer(orderId int) (int, error) {
	return o.repo.GetOrderCustomer(orderId)
}

func NewOrderService(repo *db.Repository) *OrderService {
	return &OrderService{repo: repo}
}
