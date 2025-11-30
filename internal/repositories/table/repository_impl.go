package tablerepository

import (
	tablemodels "aplikasi_restoran/internal/models/tables"

	"gorm.io/gorm"
)

type tableRepo struct { // struct implementasi interface
	db *gorm.DB // koneksi gorm
}

func NewTableRepository(db *gorm.DB) TableRepository { // constructor
	return &tableRepo{db: db} // buat instance repo
}

func (r *tableRepo) Create(user *tablemodels.Table) error { // implement create
	return r.db.Create(user).Error // insert ke database
}

func (r *tableRepo) FindById(id uint) (*tablemodels.Table, error) {
	var table tablemodels.Table
	err := r.db.First(&table, id).Error

return &table, err
}

func (r *tableRepo) Update(table *tablemodels.Table) error {
return r.db.Save(table).Error
} 

func (r *tableRepo) Delete(table *tablemodels.Table) error {
return r.db.Delete(table).Error
} 
