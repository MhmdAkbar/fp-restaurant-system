package orderdetailservice

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
	menurepo "aplikasi_restoran/internal/repositories/menu"
	orderrepo "aplikasi_restoran/internal/repositories/order"
	orderdetailrepo "aplikasi_restoran/internal/repositories/order_detail"
	orderservice "aplikasi_restoran/internal/services/order"
)

type orderDetailService struct {
	repo         orderdetailrepo.OrderDetailRepository
	menuRepo     menurepo.MenuRepository
	orderRepo    orderrepo.OrderRepository
	orderService orderservice.OrderService
}

func NewOrderDetailService(r orderdetailrepo.OrderDetailRepository, mr menurepo.MenuRepository, or orderrepo.OrderRepository, os orderservice.OrderService) OrderDetailService {
	return &orderDetailService{
		repo:         r,
		menuRepo:     mr,
		orderRepo:    or,
		orderService: os,
	}
}
func (s *orderDetailService) AddDetail(input dto.AddOrderDetailRequest) (*models.OrderDetail, error) {

	menu, err := s.menuRepo.GetById(input.MenuId)
	if err != nil {
		return nil, err
	}

	detail := &models.OrderDetail{
		OrderId:  input.OrderId,
		MenuId:   input.MenuId,
		Quantity: input.Quantity,
		Price:    menu.Price,
		Total:    menu.Price * float64(input.Quantity),
	}

	// Save detail
	newDetail, err := s.repo.Create(detail)
	if err != nil {
		return nil, err
	}

	// Recalculate order totals after save
	if err := s.orderService.RecalculateOrder(detail.OrderId); err != nil {
		return nil, err
	}

	return newDetail, nil
}

func (s *orderDetailService) UpdateDetail(detailID uint, qty uint) (*models.OrderDetail, error) {
	// Ambil detail lama
	detail, err := s.repo.GetByID(detailID)
	if err != nil {
		return nil, err
	}

	// Jika qty 0 -> controller yang handle delete
	if qty == 0 {
		return nil, nil
	}

	// Ambil harga menu terbaru
	menu, err := s.menuRepo.GetById(detail.MenuId)
	if err != nil {
		return nil, err
	}

	// Update qty & total
	detail.Quantity = qty
	detail.Price = menu.Price
	detail.Total = menu.Price * float64(qty)

	if err := s.repo.Update(detail); err != nil {
		return nil, err
	}

	// Recalculate setelah update
	if err := s.orderService.RecalculateOrder(detail.OrderId); err != nil {
		return nil, err
	}

	// Return detail terbaru + preload menu
	updated, err := s.repo.GetByID(detail.ID)
	if err != nil {
		return nil, err
	}
	return updated, nil
}


func (s *orderDetailService) DeleteDetail(detailID uint) error {
	detail, err := s.repo.GetByID(detailID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(detailID); err != nil {
		return err
	}

	return s.orderService.RecalculateOrder(detail.OrderId)
}
