package menuservice

import (
	"aplikasi_restoran/internal/models"
	menurepository "aplikasi_restoran/internal/repositories/menu"
)

type menuService struct {
	repo menurepository.MenuRepository
}

func NewMenuService(repo menurepository.MenuRepository) MenuService {
	return &menuService{repo: repo}
}

func (s *menuService) AddMenu(name string, category models.MenuCategory, price float64, desc string, img string, avail bool) (*models.Menu, error) {
	menu := models.Menu{Name: name, Category: category, Price: price, Description: desc, ImageURL: img, IsAvailable: avail}
	return &menu, s.repo.Create(&menu)
}

func (s *menuService) GetAllMenu() ([]models.Menu, error) {
	return s.repo.GetAll()
}

func (s *menuService) GetMenu(id uint) (*models.Menu, error) {
	return s.repo.GetById(id)
}

func (s *menuService) UpdateMenu(id uint, category models.MenuCategory, name string, price float64, desc string, img string, avail bool) (*models.Menu, error) {
	menu, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	menu.Name = name
	menu.Category = category
	menu.Price = price
	menu.Description = desc
	menu.ImageURL = img
	menu.IsAvailable = avail

	return menu, s.repo.Update(menu)
}

func (s *menuService) DeleteMenu(id uint) error {
	return s.repo.Delete(id)
}
