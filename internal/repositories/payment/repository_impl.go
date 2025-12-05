package paymentrepository

import (
	"aplikasi_restoran/internal/models"
	"gorm.io/gorm"
)

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepo{db: db}
}

func (r *paymentRepo) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepo) FindOrderById(orderId uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, orderId).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *paymentRepo) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepo) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}
