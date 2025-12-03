package orderservice

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
	menurepo "aplikasi_restoran/internal/repositories/menu"
	orderrepo "aplikasi_restoran/internal/repositories/order"
	orderdetailrepo "aplikasi_restoran/internal/repositories/order_detail"
	"fmt"
)

type orderService struct {
	orderRepo       orderrepo.OrderRepository
	orderDetailRepo orderdetailrepo.OrderDetailRepository
	menuRepo        menurepo.MenuRepository
}

func NewOrderService(or orderrepo.OrderRepository, od orderdetailrepo.OrderDetailRepository, mr menurepo.MenuRepository) OrderService {
	return &orderService{
		orderRepo:       or,
		orderDetailRepo: od,
		menuRepo:        mr,
	}
}
func (s *orderService) CreateOrder(input dto.CustomerCreateOrderRequest) (*models.Order, error) {
	order := models.Order{
		TableId:   input.TableId,
		Name:      input.Name,
		OrderType: input.OrderType,
		Status:    input.Status,
	}
	fmt.Println("=============")
	fmt.Println(order)
	fmt.Println("=============")
	newOrder, err := s.orderRepo.Create(&order)
	if err != nil {
		return nil, err
	}


	// ambil ulang order lengkap dengan detail
	fullOrder, err := s.orderRepo.GetWithDetails(newOrder.ID)
	if err != nil {
		return nil, err
	}
	return fullOrder, nil
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


