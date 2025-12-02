package menuservice

import "aplikasi_restoran/internal/models"

type MenuService interface {
	AddMenu(name string, category models.MenuCategory, price float64, description string, imageUrl string, isAvaliable bool) (*models.Menu, error)
	GetAllMenu() ([]models.Menu, error)
	GetMenu(id uint) (*models.Menu, error)
	UpdateMenu(id uint, category models.MenuCategory, name string, price float64, desc string, img string, avail bool) (*models.Menu, error)
	DeleteMenu(id uint) error
}
