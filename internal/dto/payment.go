package dto

import "aplikasi_restoran/internal/models"

type CreatePaymentRequest struct {
    Method models.PaymentMethod `json:"method" binding:"required"`
}
