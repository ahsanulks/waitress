package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type OrderUsecase struct {
	orderRepository OrderRepository
}

type OrderRepository interface {
	Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error)
}

func NewOrderUsecase(orderRepository OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepository: orderRepository}
}
