package tablerepository

import (
	"aplikasi_restoran/internal/models"
)

type TableRepository interface {
	Create(table *models.Table) error
	FindAll() ([]models.Table, error)
	FindById(id uint) (*models.Table, error)
	Update(table *models.Table) error
}
