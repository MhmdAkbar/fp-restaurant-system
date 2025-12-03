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
func (r *orderRepo) Create(order *models.Order) (*models.Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepo) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, id).Error
	return &order, err
}

func (r *orderRepo) GetWithDetails(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Details.Menu").First(&order, id).Error
	return &order, err
}

func (r *orderRepo) FindWithDetails(id uint) (*models.Order, error) {
    var order models.Order
    err := r.db.Preload("Details").First(&order, id).Error
    return &order, err
}

func (r *orderRepo) Update(order *models.Order) error {
    return r.db.Save(order).Error
}



