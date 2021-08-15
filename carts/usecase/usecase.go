package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type CartUsecase struct {
	cartRepository CartRepository
}

// CartRepository is all dependency method that needed by cart usecase.
type CartRepository interface {
	FindByUserID(ctx context.Context, userID int) (domain.Cart, error)
	Create(ctx context.Context, cart *domain.Cart) error
	AddItem(ctx context.Context, cartItem *domain.CartItem) error
}

// NewCartUsecase to create new cart usecase.
func NewCartUsecase(cartRepository CartRepository) *CartUsecase {
	return &CartUsecase{cartRepository: cartRepository}
}
