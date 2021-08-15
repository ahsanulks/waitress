package repository_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/carts/repository"
	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
)

func TestCartRepository_FindByUserID(t *testing.T) {
	var (
		db         = reltest.New()
		cartResult = domain.Cart{
			ID:     1,
			UserID: 2,
		}
		cartItems = []domain.CartItem{
			{
				CartID:    1,
				ProductID: 2,
			},
		}
		product = domain.Product{
			ID:   2,
			Name: "product test",
		}
		cartItemWithProduct = []domain.CartItem{
			{
				CartID:    1,
				ProductID: 2,
				Product:   &product,
			},
		}
		cartWithItemsAndProduct = domain.Cart{
			ID:        1,
			UserID:    2,
			CartItems: cartItemWithProduct,
		}
	)
	tests := []struct {
		name     string
		userID   int
		want     domain.Cart
		wantErr  bool
		mockFunc func()
	}{
		{
			name:    "found without preload product",
			userID:  2,
			want:    cartResult,
			wantErr: false,
			mockFunc: func() {
				db.ExpectFind(rel.Eq("user_id", 2)).Result(cartResult)
				db.ExpectPreload("cart_items", rel.Eq("purchased", false))
			},
		},
		{
			name:    "found with preload product",
			userID:  2,
			want:    cartWithItemsAndProduct,
			wantErr: false,
			mockFunc: func() {
				db.ExpectFind(rel.Eq("user_id", 2)).Result(cartResult)
				db.ExpectPreload("cart_items", rel.Eq("purchased", false)).Result(cartItems)
				db.ExpectPreload("product").Result(product)
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
			got, err := repo.FindByUserID(ctx, tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepository.FindByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepository.FindByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
