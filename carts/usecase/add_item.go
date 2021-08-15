package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

// AddItem insert cart item based on cart_id and product_id.
func (cu CartUsecase) AddItem(ctx context.Context, cartItemParams domain.CartItemParams) (domain.CartItem, error) {
	var cartItem domain.CartItem
	validate := validator.New()
	if err := validate.Struct(cartItemParams); err != nil {
		return cartItem, err
	}

	copier.Copy(&cartItem, cartItemParams)
	err := cu.cartRepository.AddItem(ctx, &cartItem)
	return cartItem, err
}
