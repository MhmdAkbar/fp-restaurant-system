package paymentrepository

import "aplikasi_restoran/internal/models"

type PaymentRepository interface {
	Create(payment *models.Payment) error
	FindOrderById(orderId uint) (*models.Order, error)
	UpdatePayment(payment *models.Payment) error
	UpdateOrder(order *models.Order) error
}
