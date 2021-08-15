package repository_test

import (
	"context"
	"testing"

	"github.com/ahsanulks/waitress/carts/repository"
	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
)

func TestCartRepository_AddItem(t *testing.T) {
	var (
		db       = reltest.New()
		cartItem = domain.CartItem{
			CartID:    1,
			ProductID: 2,
			Quantity:  3,
		}
		cart = domain.Cart{
			ID:     1,
			UserID: 3,
		}
		productNotEnough = domain.Product{
			ID:    2,
			Stock: 1,
		}
		product = domain.Product{
			ID:    2,
			Stock: 5,
		}
	)
	tests := []struct {
		name     string
		cartItem *domain.CartItem
		wantErr  bool
		mockFunc func()
	}{
		{
			name:     "cannot find cart",
			cartItem: &cartItem,
			wantErr:  true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					repo.ExpectFind(rel.Eq("id", 1)).NotFound()
				})
			},
		},
		{
			name:     "cannot find product",
			cartItem: &cartItem,
			wantErr:  true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					repo.ExpectFind(rel.Eq("id", 1)).Result(cart)
					repo.ExpectFind(rel.Eq("id", 2), rel.ForUpdate()).NotFound()
				})
			},
		},
		{
			name:     "still have a same product that not purchased",
			cartItem: &cartItem,
			wantErr:  true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					repo.ExpectFind(rel.Eq("id", 1)).Result(cart)
					repo.ExpectFind(rel.Eq("id", 2), rel.ForUpdate()).Result(productNotEnough)
					repo.ExpectFind(rel.Eq("purchased", false).AndEq("cart_id", cartItem.CartID).AndEq("product_id", cartItem.ProductID)).Result(cartItem)
				})
			},
		},
		{
			name:     "stock not enough",
			cartItem: &cartItem,
			wantErr:  true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					repo.ExpectFind(rel.Eq("id", 1)).Result(cart)
					repo.ExpectFind(rel.Eq("id", 2), rel.ForUpdate()).Result(productNotEnough)
					repo.ExpectFind(rel.Eq("purchased", false).AndEq("cart_id", cartItem.CartID).AndEq("product_id", cartItem.ProductID)).NotFound()
				})
			},
		},
		{
			name:     "success fully create",
			cartItem: &cartItem,
			wantErr:  false,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					repo.ExpectFind(rel.Eq("id", 1)).Result(cart)
					repo.ExpectFind(rel.Eq("id", 2), rel.ForUpdate()).Result(product)
					repo.ExpectFind(rel.Eq("purchased", false).AndEq("cart_id", cartItem.CartID).AndEq("product_id", cartItem.ProductID)).NotFound()
					repo.ExpectInsert().For(&cartItem)
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				repo = repository.NewCartRepository(db)
				ctx  = context.TODO()
			)
			tt.mockFunc()
			if err := repo.AddItem(ctx, tt.cartItem); (err != nil) != tt.wantErr {
				t.Errorf("CartRepository.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
