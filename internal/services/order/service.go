package orderservice

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
)

type OrderService interface {
	CreateOrder(input dto.CustomerCreateOrderRequest) (*models.Order, error)
	GetOrder(id uint) (*models.Order, error)
	RecalculateOrder(orderID uint) error
	UpdateStatus(id uint, status models.OrderStatus) error
}
