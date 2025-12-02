package orderrepository

import (
	"aplikasi_restoran/internal/models"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) GetByID(id uint) (*models.Order, error) {
	var o models.Order
	err := r.db.Preload("OrderDetails").First(&o, id).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *orderRepo) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepo) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
