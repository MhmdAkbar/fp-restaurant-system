package tableservice

import (
	"aplikasi_restoran/internal/models"
	tablerepository "aplikasi_restoran/internal/repositories/table"
	"fmt"
)


type tableService struct {
	repo tablerepository.TableRepository
}

func NewTableService(repo tablerepository.TableRepository) TableService {
	return &tableService{repo}
}

func (s *tableService) AddTable(qrCode string) (*models.Table, error) {
	table := models.Table{
		QrCode: qrCode,
	}

	// Step 1: Create
	if err := s.repo.Create(&table); err != nil {
		return nil, err
	}

	// Step 2: Generate TableCode
	table.TableCode = fmt.Sprintf("TBL-%03d", table.ID)

	// Step 3: Update lagi
	if err := s.repo.Update(&table); err != nil {
		return nil, err
	}

	return &table, nil
}

func (s *tableService) GetTable(id uint) (*models.Table, error) {
	return s.repo.FindById(id)
}

func (s *tableService) GetAll() ([]models.Table, error) {
	return s.repo.FindAll()
}

func (s *tableService) UpdateTable(id uint, qrCode string) (*models.Table, error) {
	table, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	table.QrCode = qrCode

	return table, s.repo.Update(table)
}
func (s *tableService) UpdateStatus(id uint, status bool) (*models.Table, error) {
    table, err := s.repo.FindById(id)
    if err != nil {
        return nil, err
    }

    table.Status = status

    // WAJIB SIMPAN KE DB!!
    if err := s.repo.Update(table); err != nil {
        return nil, err
    }

    return table, nil
}

