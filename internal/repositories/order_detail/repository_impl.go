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
func (r *orderDetailRepo) Create(detail *models.OrderDetail) (*models.OrderDetail, error) {
	if err := r.db.Create(detail).Error; err != nil {
		return nil, err
	}
	return detail, nil
}


func (r *orderDetailRepo) GetByOrderID(orderID uint) ([]models.OrderDetail, error) {
	var details []models.OrderDetail
	err := r.db.Where("order_id = ?", orderID).Preload("Menu").Find(&details).Error
	return details, err
}

func (r *orderDetailRepo) GetByID(id uint) (*models.OrderDetail, error) {
	var d models.OrderDetail
	if err := r.db.First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *orderDetailRepo) DeleteByOrderID(orderID uint) error {
	return r.db.Where("order_id = ?", orderID).Delete(&models.OrderDetail{}).Error
}

func (r *orderDetailRepo) Update(detail *models.OrderDetail) error {
	return r.db.Save(detail).Error
}

func (r *orderDetailRepo) Delete(detailID uint) error {
	return r.db.Delete(&models.OrderDetail{}, detailID).Error
}
