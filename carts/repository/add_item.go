package repository

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

// AddItem for add cart_item based on cart_id and product that have enough stock.
func (cr CartRepository) AddItem(ctx context.Context, cartItem *domain.CartItem) error {
	return cr.db.Transaction(ctx, func(ctx context.Context) error {
		var (
			cart          domain.Cart
			product       domain.Product
			checkCartItem domain.CartItem
		)
		// check that cart relation is present
		if err := cr.db.Find(ctx, &cart, rel.Eq("id", cartItem.CartID)); err != nil {
			return errors.New("cart not found")
		}

		// check that product relation is present
		// lock the product
		if err := cr.db.Find(ctx, &product, rel.Eq("id", cartItem.ProductID), rel.ForUpdate()); err != nil {
			return errors.New("product not found")
		}

		// check that doens't have same product id on cart item is still not purchased
		if err := cr.db.Find(ctx, &checkCartItem, rel.Eq("purchased", false).AndEq("cart_id", cartItem.CartID).AndEq("product_id", cartItem.ProductID)); err == nil && checkCartItem.ID != 0 {
			return errors.New("cant create new card item with same product id")
		}

		// check that product stock >= quantity on cart item requested
		if cartItem.Quantity > product.Stock {
			return errors.New("stock not enough")
		}

		return cr.db.Insert(ctx, cartItem)
	})
}
