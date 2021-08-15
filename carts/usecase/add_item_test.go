package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/carts/usecase"
	"github.com/ahsanulks/waitress/domain"
)

func TestCartUsecase_AddItem(t *testing.T) {
	tests := []struct {
		name           string
		cartItemParams domain.CartItemParams
		want           domain.CartItem
		wantErr        bool
	}{
		{
			name: "error when validate data",
			cartItemParams: domain.CartItemParams{
				CartID: 2,
			},
			want:    domain.CartItem{},
			wantErr: true,
		},
		{
			name: "success add item",
			cartItemParams: domain.CartItemParams{
				CartID:    2,
				ProductID: 1,
				Quantity:  2,
			},
			want: domain.CartItem{
				CartID:    2,
				ProductID: 1,
				Quantity:  2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				cu  = usecase.NewCartUsecase(fakeCartRepository{})
				ctx = context.TODO()
			)
			got, err := cu.AddItem(ctx, tt.cartItemParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
