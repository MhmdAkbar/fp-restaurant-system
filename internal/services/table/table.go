package tableservice

import tablemodels "aplikasi_restoran/internal/models/tables"


type TableService interface {
	AddTable(qrCode string, status bool) (*tablemodels.Table, error)
	GetTable(id uint) (*tablemodels.Table, error)
	GetAll() ([]tablemodels.Table, error)
	UpdateTable(id uint, qrCode string, status bool) (*tablemodels.Table, error)
	UpdateStatus(id uint, status bool) (*tablemodels.Table, error)
}