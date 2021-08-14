package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/domain"
	"github.com/ahsanulks/waitress/products/usecase"
)

func TestProductUsecase_FindAll(t *testing.T) {
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.Product
		wantErr bool
	}{
		{
			name: "found",
			args: args{limit: 10, offset: 0},
			want: []domain.Product{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			wantErr: false,
		},
		{
			name:    "not found",
			args:    args{limit: 1, offset: 100},
			want:    []domain.Product{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				pu  = usecase.NewProductUsecase(fakeProductRepository{})
				ctx = context.TODO()
			)
			got, err := pu.FindAll(ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductUsecase.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductUsecase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
