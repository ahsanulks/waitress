package repository

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

// FindByUserID is for find cart items by user_id that not purchased yet
func (cr CartRepository) FindByUserID(ctx context.Context, userID int) (domain.Cart, error) {
	var cart domain.Cart
	err := cr.db.Find(ctx, &cart, rel.Eq("user_id", userID))
	if err != nil {
		return domain.Cart{}, errors.New("cart not found")
	}
	// load cart_items
	cr.db.Preload(ctx, &cart, "cart_items", rel.Eq("purchased", false))
	if len(cart.CartItems) > 0 {
		cr.db.Preload(ctx, &cart.CartItems, "product")
	}
	return cart, err
}
