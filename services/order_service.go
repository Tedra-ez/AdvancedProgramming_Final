package services

import (
	"context"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"github.com/Tedra-ez/AdvancedProgramming_Final/repository"
)

type OrderService struct {
	orderRepo repository.OrderStore
}

func NewOrderService(orderRepo repository.OrderStore) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) Create(ctx context.Context, req *models.CreateOrderRequest) (*models.Order, error) {
	var subtotal float64
	items := make([]models.OrderItem, 0, len(req.Items))
	for _, it := range req.Items {
		lineTotal := it.UnitPrice * float64(it.Quantity)
		subtotal += lineTotal
		items = append(items, models.OrderItem{
			ProductID:     it.ProductID,
			ProductName:   it.ProductName,
			SelectedSize:  it.SelectedSize,
			SelectedColor: it.SelectedColor,
			Quantity:      it.Quantity,
			UnitPrice:     it.UnitPrice,
			LineTotal:     lineTotal,
		})
	}
	total := subtotal
	order := &models.Order{
		UserID:          req.UserID,
		Status:          "pending",
		PaymentMethod:   req.PaymentMethod,
		DeliveryMethod:  req.DeliveryMethod,
		DeliveryAddress: req.DeliveryAddress,
		Comment:         req.Comment,
		Subtotal:        subtotal,
		DeliveryFee:     0,
		Total:           total,
		Items:           items,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := s.orderRepo.Save(ctx, order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) ListByUser(ctx context.Context, userID string) ([]*models.Order, error) {
	return s.orderRepo.FindByUser(ctx, userID)
}

func (s *OrderService) UpdateStatus(ctx context.Context, orderID, status string) error {
	return s.orderRepo.UpdateStatus(ctx, orderID, status)
}

func (s *OrderService) GetByID(ctx context.Context, orderID string) (*models.Order, error) {
	return s.orderRepo.FindByID(ctx, orderID)
}
