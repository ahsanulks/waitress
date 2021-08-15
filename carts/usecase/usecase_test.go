package usecase_test

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
)

type fakeCartRepository struct{}

func (fcr fakeCartRepository) FindByUserID(ctx context.Context, userID int) (domain.Cart, error) {
	if userID == 2 || userID == 4 {
		return domain.Cart{}, errors.New("not found")
	}
	return domain.Cart{
		ID:     1,
		UserID: userID,
	}, nil
}

func (fcr fakeCartRepository) Create(ctx context.Context, cart *domain.Cart) error {
	if cart.UserID == 2 {
		cart.ID = 3
		return nil
	}
	return errors.New("error")
}

func (fcr fakeCartRepository) AddItem(ctx context.Context, cartItem *domain.CartItem) error {
	return nil
}
