package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

func (cr CartRepository) FindByUserID(ctx context.Context, userID int) (domain.Cart, error) {
	var cart domain.Cart
	err := cr.db.Find(ctx, &cart, rel.Eq("user_id", userID))
	// load cart_items
	cr.db.Preload(ctx, &cart, "cart_items", rel.Eq("purchased", false))
	if len(cart.CartItems) > 0 {
		cr.db.FindAll(ctx, &cart.CartItems, rel.Preload("product"))
	}
	return cart, err
}
