package menurepository

import (
	"aplikasi_restoran/internal/models"
	"errors"

	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepo{db: db}
}

func (r *menuRepo) Create(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepo) GetAll() ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.Find(&menus).Error
	return menus, err
}

func (r *menuRepo) GetById(id uint) (*models.Menu, error) {
	var menu models.Menu
	if err := r.db.First(&menu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("menu not found")
		}
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepo) Update(menu *models.Menu) error {
	return r.db.Save(menu).Error
}

func (r *menuRepo) Delete(id uint) error {
	return r.db.Delete(&models.Menu{}, id).Error
}
