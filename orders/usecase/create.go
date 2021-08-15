package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-playground/validator/v10"
)

func (ou OrderUsecase) Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error) {
	validate := validator.New()
	if err := validate.Struct(orderPrams); err != nil {
		return domain.Order{}, err
	}
	return ou.orderRepository.Create(ctx, orderPrams)
}
