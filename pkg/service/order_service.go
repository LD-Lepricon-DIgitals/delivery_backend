package service

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
)

type OrderService struct {
	repo *db.Repository
}

func (o OrderService) GetOrderDetails(orderId int) (models.OrderDetails, error) {
	return o.repo.GetOrderDetails(orderId)
}

func (o OrderService) CreateOrder(order models.CreateOrder) error {
	return o.repo.CreateOrder(order)
}

func (o OrderService) GetOrders(workerId int) ([]models.OrderInfo, error) {
	return o.repo.GetOrders(workerId)
}

func (o OrderService) FinishOrder(orderId, workerId int) error {
	return o.repo.FinishOrder(orderId, workerId)
}

func (o OrderService) StartOrder(orderId int, workerId int) error {
	return o.repo.StartOrder(orderId, workerId)
}

func NewOrderService(repo *db.Repository) *OrderService {
	return &OrderService{repo: repo}
}
