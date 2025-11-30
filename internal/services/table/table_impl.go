package tableservice

import (
	"fmt"
	tablemodels "aplikasi_restoran/internal/models/tables"
	tablerepository "aplikasi_restoran/internal/repositories/table"
)


type tableService struct {
	repo tablerepository.TableRepository
}

func NewTableService(repo tablerepository.TableRepository) TableService {
	return &tableService{repo}
}

func (s *tableService) AddTable(qrCode string, status bool) (*tablemodels.Table, error) {
	table := tablemodels.Table{
		QrCode: qrCode,
		Status: status,
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

func (s *tableService) GetTable(id uint) (*tablemodels.Table, error) {
	return s.repo.FindById(id)
}

func (s *tableService) GetAll() ([]tablemodels.Table, error) {
	return s.repo.FindAll()
}

func (s *tableService) UpdateTable(id uint, qrCode string, status bool) (*tablemodels.Table, error) {
	table, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	table.QrCode = qrCode
	table.Status = status

	return table, s.repo.Update(table)
}

func (s *tableService) UpdateStatus(id uint, status bool) (*tablemodels.Table, error) {
	table, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	table.Status = status
	return table, s.repo.Update(table)
}
