package tablerepository

import tablemodels "aplikasi_restoran/internal/models/tables"

type TableRepository interface {
	Create(table *tablemodels.Table) error
	FindAll() ([]tablemodels.Table, error)
	FindById(id uint) (*tablemodels.Table, error)
	Update(table *tablemodels.Table) error
}