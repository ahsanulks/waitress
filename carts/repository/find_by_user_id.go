package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

// FindByUserID is for find cart items by user_id that not purchased yet
func (cr CartRepository) FindByUserID(ctx context.Context, userID int) (domain.Cart, error) {
	var cart domain.Cart
	err := cr.db.Find(ctx, &cart, rel.Eq("user_id", userID))
	// load cart_items
	cr.db.Preload(ctx, &cart, "cart_items", rel.Eq("purchased", false))
	cartItems := cart.CartItems
	if len(cartItems) > 0 {
		cr.db.FindAll(ctx, &cartItems, rel.Preload("product"))
	}
	return cart, err
}
