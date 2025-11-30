package tableservice

import (
	tablemodels "aplikasi_restoran/internal/models/tables"
	repositories "aplikasi_restoran/internal/repositories/table"
)

type tableService struct {
	repo repositories.TableRepository
}

func NewTableService(repo repositories.TableRepository) TableService {
	return &tableService{repo}
}

func (s *tableService) AddTable(tableNum uint, qrCode string, avaliable bool) (*tablemodels.Table, error) {
	table := tablemodels.Table{
		TableNumber: tableNum,
		QrCode:      qrCode,
		Status:      avaliable,
	}

	if err := s.repo.Create(&table); err != nil {
		return nil, err
	}

	return &table, nil
}

func (s *tableService) GetTable(id uint) (*tablemodels.Table, error) {
	return s.repo.FindById(id)
}

func (s *tableService) UpdateTable(id, tableNum uint, qrCode string, avaliable bool) (*tablemodels.Table, error) {
	table, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	table.TableNumber = tableNum
	table.QrCode = qrCode
	table.Status = avaliable

	if err := s.repo.Update(table); err != nil {
		return nil, err
	}

	return table, nil
}

func (s *tableService) UpdateStatus(id uint, avaliable bool) (*tablemodels.Table, error) {

	table, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	table.Status = avaliable

	if err := s.repo.Update(table); err != nil {
		return nil, err
	}

	return table, nil
}

func (s *tableService) DeleteTable(id uint) error {

	table, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(table); err != nil {
		return err
	}
	return nil
}
