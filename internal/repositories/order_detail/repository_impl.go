package orderdetailrepository

import (
	"aplikasi_restoran/internal/models"

	"gorm.io/gorm"
)

type orderDetailRepo struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepo{db: db}
}

func (r *orderDetailRepo) Create(detail *models.OrderDetail) error {
	return r.db.Create(detail).Error
}

func (r *orderDetailRepo) GetByOrderID(orderID uint) ([]models.OrderDetail, error) {
	var details []models.OrderDetail // variabel penampung

	err := r.db.Where("order_id = ?", orderID).Find(&details).Error
	return details, err
}

func (r *orderDetailRepo)DeleteByOrderID(orderID uint) error {
	return r.db.Where("order_id = ?", orderID).Delete(&models.OrderDetail{}).Error
}
