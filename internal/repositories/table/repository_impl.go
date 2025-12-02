package tablerepository

import (
	"aplikasi_restoran/internal/models"

	"gorm.io/gorm"
)

type tableRepo struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepo{db: db}
}

func (r *tableRepo) Create(table *models.Table) error {
	return r.db.Create(table).Error
}

func (r *tableRepo) FindAll() ([]models.Table, error) {
	var tables []models.Table
	err := r.db.Find(&tables).Error
	return tables, err
}

func (r *tableRepo) FindById(id uint) (*models.Table, error) {
	var table models.Table
	err := r.db.First(&table, id).Error
	return &table, err
}

func (r *tableRepo) Update(table *models.Table) error {
	return r.db.Save(table).Error // update seluruh field yang berubah
}
