package usecase_test

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
)

type fakeOrderRepository struct{}

func (fakeRepo fakeOrderRepository) Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error) {
	if orderPrams.BuyerID == 3 {
		return domain.Order{}, errors.New("error when create")
	}
	return domain.Order{
		ID:    2,
		State: domain.Pending,
	}, nil
}
