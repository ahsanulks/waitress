package usecase_test

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
)

type fakeProductRepository struct{}

func (fakeProductRepository) Create(ctx context.Context, product *domain.Product) error {
	if product.SellerID == 2 {
		return errors.New("error")
	}
	return nil
}

func (fakeProductRepository) FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	if limit == 1 && offset == 100 {
		return []domain.Product{}, nil
	}
	return []domain.Product{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}, nil
}
