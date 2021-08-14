package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

func (cu CartUsecase) FindOrCreate(ctx context.Context, userID int) (domain.Cart, error) {
	cart, err := cu.cartRepository.FindByUserID(ctx, userID)
	if err == nil {
		return cart, err
	}
	cart.UserID = userID
	err = cu.cartRepository.Create(ctx, &cart)
	return cart, err
}
