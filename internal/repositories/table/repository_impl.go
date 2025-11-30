package tablerepository

import (
	tablemodels "aplikasi_restoran/internal/models/tables"

	"gorm.io/gorm"
)

type tableRepo struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepo{db: db}
}

func (r *tableRepo) Create(table *tablemodels.Table) error {
	return r.db.Create(table).Error
}

func (r *tableRepo) FindAll() ([]tablemodels.Table, error) {
	var tables []tablemodels.Table
	err := r.db.Find(&tables).Error
	return tables, err
}

func (r *tableRepo) FindById(id uint) (*tablemodels.Table, error) {
	var table tablemodels.Table
	err := r.db.First(&table, id).Error
	return &table, err
}

func (r *tableRepo) Update(table *tablemodels.Table) error {
	return r.db.Save(table).Error // update seluruh field yang berubah
}
