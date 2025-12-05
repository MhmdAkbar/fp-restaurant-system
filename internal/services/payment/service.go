package paymentservice

import "aplikasi_restoran/internal/models"

type PaymentService interface {
	CreatePayment(orderId uint, cashierId uint, method models.PaymentMethod) (*models.Payment, error)
}