package orderservice

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
	menurepo "aplikasi_restoran/internal/repositories/menu"
	orderrepo "aplikasi_restoran/internal/repositories/order"
	orderdetailrepo "aplikasi_restoran/internal/repositories/order_detail"
	tablerepo "aplikasi_restoran/internal/repositories/table"
	"errors"
)

type orderService struct {
	orderRepo       orderrepo.OrderRepository
	orderDetailRepo orderdetailrepo.OrderDetailRepository
	menuRepo        menurepo.MenuRepository
	tableRepo       tablerepo.TableRepository
}

func NewOrderService(or orderrepo.OrderRepository, od orderdetailrepo.OrderDetailRepository, mr menurepo.MenuRepository, tr tablerepo.TableRepository) OrderService {
	return &orderService{or, od, mr, tr}
}

func (s *orderService) CreateOrder(input dto.CustomerCreateOrderRequest) (*models.Order, error) {
	// Cek status meja
	table, err := s.tableRepo.FindById(input.TableId)
	if err != nil {
		return nil, err
	}
	if table.Status {
		return nil, errors.New("meja masih digunakan pelanggan lain")
	}

	order := models.Order{
		TableId:   input.TableId,
		Name:      input.Name,
		OrderType: input.OrderType,
		Status:    input.Status,
	}

	newOrder, err := s.orderRepo.Create(&order)
	if err != nil {
		return nil, err
	}

	// Update status meja
	table.Status = true
	_ = s.tableRepo.Update(table)

	return s.orderRepo.GetWithDetails(newOrder.ID)
}

func (s *orderService) GetOrder(id uint) (*models.Order, error) {
	return s.orderRepo.GetWithDetails(id)
}

func (s *orderService) RecalculateOrder(orderID uint) error {
	order, err := s.orderRepo.FindWithDetails(orderID)
	if err != nil {
		return err
	}

	var subtotal float64
	for _, d := range order.Details {
		subtotal += d.Total
	}

	order.Subtotal = subtotal
	order.Tax = subtotal * 0.11
	order.ServiceCharge = subtotal * 0.05
	order.GrandTotal = subtotal + order.Tax + order.ServiceCharge - order.Discount

	return s.orderRepo.Update(order)
}

func (s *orderService) UpdateStatus(id uint, status models.OrderStatus) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	order.Status = status
	if err := s.orderRepo.Update(order); err != nil {
		return err
	}

	// Jika order selesai → meja dikosongkan
	

	return nil
}

