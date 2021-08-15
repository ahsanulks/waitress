package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/carts/usecase"
	"github.com/ahsanulks/waitress/domain"
)

func TestCartUsecase_FindOrCreate(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		want    domain.Cart
		wantErr bool
	}{
		{
			name:   "create cart",
			userID: 2,
			want: domain.Cart{
				ID:     3,
				UserID: 2,
			},
			wantErr: false,
		},
		{
			name:   "find cart",
			userID: 3,
			want: domain.Cart{
				ID:     1,
				UserID: 3,
			},
			wantErr: false,
		},
		{
			name:   "error when create cart",
			userID: 4,
			want: domain.Cart{
				UserID: 4,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				cu  = usecase.NewCartUsecase(fakeCartRepository{})
				ctx = context.TODO()
			)
			got, err := cu.FindOrCreate(ctx, tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.FindOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.FindOrCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
