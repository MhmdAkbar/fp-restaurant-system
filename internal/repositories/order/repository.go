package orderrepository

import "aplikasi_restoran/internal/models"

type OrderRepository interface { // kontrak method untuk order
	Create(order *models.Order) error       // membuat order baru
	GetByID(id uint) (*models.Order, error) // ambil order per id
	Update(order *models.Order) error       // update order
	Delete(id uint) error                   // hapus order
}
