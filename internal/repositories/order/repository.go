package orderrepository

import "aplikasi_restoran/internal/models"

type OrderRepository interface { // kontrak method untuk order
	Create(order *models.Order) (*models.Order, error) // membuat order baru
	GetByID(id uint) (*models.Order, error) // ambil order per id
	GetWithDetails(id uint) (*models.Order, error)
	FindWithDetails(id uint) (*models.Order, error)
	Update(order *models.Order) error
	// Update(order *models.Order) error       // update order
	// Delete(id uint) error                   // hapus order
	// GetActiveOrderByTable(tableId uint) (*models.Order, error)
}
