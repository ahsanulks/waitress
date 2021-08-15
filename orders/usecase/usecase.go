package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type OrderUsecase struct {
	orderRepository OrderRepository
}

// OrderRepository reference all method that needed for order usecase.
type OrderRepository interface {
	Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error)
}

// NewOrderUsecase will create new order usecase
func NewOrderUsecase(orderRepository OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepository: orderRepository}
}
