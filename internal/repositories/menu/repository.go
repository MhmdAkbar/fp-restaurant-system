package menurepository

import "aplikasi_restoran/internal/models"

type MenuRepository interface {
	Create(menu *models.Menu) error
	GetAll() ([]models.Menu, error)
	GetById(id uint) (*models.Menu, error)
	Update(menu *models.Menu) error
	Delete(id uint) error
}