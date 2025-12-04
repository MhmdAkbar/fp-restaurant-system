package tableservice

import "aplikasi_restoran/internal/models"

type TableService interface {
	AddTable(qrCode string) (*models.Table, error)
	GetTable(id uint) (*models.Table, error)
	GetAll() ([]models.Table, error)
	UpdateTable(id uint, qrCode string) (*models.Table, error)
	UpdateStatus(id uint, status bool) (*models.Table, error)
}
