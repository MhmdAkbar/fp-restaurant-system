package tableservice

import tablemodels "aplikasi_restoran/internal/models/tables"

type TableService interface {
	AddTable(tableNum uint, qrCode string, avaliable bool) (*tablemodels.Table, error)
	GetTable(id uint) (*tablemodels.Table, error)
	UpdateTable(id ,tableNum uint, qrCode string, avaliable bool) (*tablemodels.Table, error)
	UpdateStatus(id uint,avaliable bool)(*tablemodels.Table, error)
	DeleteTable(id uint) error
}