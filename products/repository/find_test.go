package repository_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/domain"
	"github.com/ahsanulks/waitress/products/repository"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
)

func Test_findProduct_FindAll(t *testing.T) {
	var (
		productResult = []domain.Product{{ID: 1}}
		db            = reltest.New()
	)
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name     string
		args     args
		want     []domain.Product
		wantErr  bool
		mockFunc func()
	}{
		{
			name:    "success found",
			args:    args{limit: 1, offset: 0},
			want:    productResult,
			wantErr: false,
			mockFunc: func() {
				db.ExpectFindAll(rel.Offset(0), rel.Limit(1)).Result(productResult)
			},
		},
		{
			name:    "error when connect to db",
			args:    args{limit: 2, offset: 0},
			want:    []domain.Product{{}},
			wantErr: true,
			mockFunc: func() {
				db.ExpectFindAll(rel.Offset(0), rel.Limit(2)).Error(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				repo = repository.NewProductRepository(db)
				ctx  = context.TODO()
			)
			tt.mockFunc()
			got, err := repo.FindAll(ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("findProduct.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findProduct.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
