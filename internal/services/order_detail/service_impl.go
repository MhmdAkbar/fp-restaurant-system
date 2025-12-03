package orderdetailservice

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
	menurepo "aplikasi_restoran/internal/repositories/menu"
	orderrepo "aplikasi_restoran/internal/repositories/order"
	orderservice "aplikasi_restoran/internal/services/order"
	orderdetailrepo "aplikasi_restoran/internal/repositories/order_detail"
)

type orderDetailService struct {
	repo      orderdetailrepo.OrderDetailRepository
	menuRepo  menurepo.MenuRepository
	orderRepo orderrepo.OrderRepository
	 orderService orderservice.OrderService
}

func NewOrderDetailService(r orderdetailrepo.OrderDetailRepository, mr menurepo.MenuRepository, or orderrepo.OrderRepository, os orderservice.OrderService) OrderDetailService {
	return &orderDetailService{
		repo:      r,
		menuRepo:  mr,
		orderRepo: or,
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


// func (s *orderDetailService) UpdateDetail(detailID uint, qty uint) (*models.OrderDetail, error) {
// 	detail, err := s.repo.GetByID(detailID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	menu, err := s.menuRepo.GetById(detail.MenuId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	oldQty := int(detail.Quantity)
// 	newQty := int(qty)
// 	delta := newQty - oldQty

// 	// transaction
// 	db := s.repo.DB()
// 	tx := db.Begin()
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	if delta > 0 {
// 		if menu.Stock < delta {
// 			tx.Rollback()
// 			return nil, errors.New("not enough stock to increase quantity")
// 		}
// 		if err := tx.Model(&models.Menu{}).Where("id = ?", menu.ID).Update("stock", menu.Stock-delta).Error; err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}
// 	} else if delta < 0 {
// 		if err := tx.Model(&models.Menu{}).Where("id = ?", menu.ID).Update("stock", menu.Stock+(-delta)).Error; err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}
// 	}

// 	detail.Quantity = qty
// 	detail.Total = detail.Price * float64(qty)
// 	if err := tx.Save(detail).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	// recalc order totals
// 	order, err := s.orderRepo.GetByID(detail.OrderId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	details, _ := s.repo.GetByOrderID(order.ID)
// 	var subtotal float64
// 	for _, dd := range details {
// 		subtotal += dd.Total
// 	}
// 	order.Subtotal = subtotal
// 	order.GrandTotal = subtotal + order.Tax + order.ServiceCharge - order.Discount
// 	if err := s.orderRepo.Update(order); err != nil {
// 		return nil, err
// 	}

// 	return detail, nil
// }

// func (s *orderDetailService) DeleteDetail(detailID uint) error {
// 	detail, err := s.repo.GetByID(detailID)
// 	if err != nil {
// 		return err
// 	}

// 	// transaction: restore stock, delete detail, recalc totals
// 	db := s.repo.DB()
// 	tx := db.Begin()
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	// restore stock
// 	menu, err := s.menuRepo.GetById(detail.MenuId)
// 	if err == nil {
// 		if err := tx.Model(&models.Menu{}).Where("id = ?", menu.ID).Update("stock", menu.Stock+int(detail.Quantity)).Error; err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}

// 	// delete detail
// 	if err := tx.Delete(&models.OrderDetail{}, detailID).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// recalc
// 	order, err := s.orderRepo.GetByID(detail.OrderId)
// 	if err != nil {
// 		// if order not found, ignore
// 		return nil
// 	}
// 	details, _ := s.repo.GetByOrderID(order.ID)
// 	var subtotal float64
// 	for _, dd := range details {
// 		subtotal += dd.Total
// 	}
// 	order.Subtotal = subtotal
// 	order.GrandTotal = subtotal + order.Tax + order.ServiceCharge - order.Discount
// 	return s.orderRepo.Update(order)
// }

// func (s *orderDetailService) GetOrderDetails(orderID uint) ([]models.OrderDetail, error) {
// 	return s.repo.GetByOrderID(orderID)
// }
