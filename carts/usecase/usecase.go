package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type CartUsecase struct {
	cartRepository CartRepository
}

type CartRepository interface {
	FindByUserID(ctx context.Context, userID int) (domain.Cart, error)
	Create(ctx context.Context, cart *domain.Cart) error
	AddItem(ctx context.Context, cartItem *domain.CartItem) error
}

func NewCartUsecase(cartRepository CartRepository) *CartUsecase {
	return &CartUsecase{cartRepository: cartRepository}
}
