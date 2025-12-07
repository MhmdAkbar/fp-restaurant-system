package paymentservice

import (
	"aplikasi_restoran/internal/models"
	paymentrepository "aplikasi_restoran/internal/repositories/payment"
	tablerepository "aplikasi_restoran/internal/repositories/table"
	"errors"
)

type paymentService struct {
	repo paymentrepository.PaymentRepository
	tableRepo tablerepository.TableRepository
}

func NewPaymentService(pr paymentrepository.PaymentRepository, tr tablerepository.TableRepository) PaymentService {
	return &paymentService{pr,tr}
}

func (s *paymentService) CreatePayment(orderId uint, cashierId uint, method models.PaymentMethod) (*models.Payment, error) {
	// Ambil order
	order, err := s.repo.FindOrderById(orderId)
	if err != nil {
		return nil, err
	}

	// Siapkan pointer untuk orderID
	orderID := order.ID
	if order.Status == models.OrderDone{
		return nil , errors.New("this order has been paid")
	}

	// Buat pembayaran langsung lunas
	payment := models.Payment{
		OrderID:       &orderID,
		CashierID:     cashierId,
		PaymentMethod: method,
		AmountPaid:    order.GrandTotal, // ⬅ AUTO terisi
		PaymentDone:   true,             // ⬅ Selalu lunas
	}

	// Simpan payment
	if err := s.repo.Create(&payment); err != nil {
		return nil, err
	}

	// Update order: selesai disajikan
	order.Status = models.OrderDone
	if err := s.repo.UpdateOrder(order); err != nil {
		return nil, err
	}
	if models.Done {
		table, _ := s.tableRepo.FindById(order.TableId)
		table.Status = false
		_ = s.tableRepo.Update(table)
	}

	return &payment, nil
}
