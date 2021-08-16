package usecase

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-playground/validator/v10"
)

// Create will check all need data and passing to repository to create order
func (ou OrderUsecase) Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error) {
	validate := validator.New()
	if err := validate.Struct(orderPrams); err != nil {
		return domain.Order{}, errors.New("parameters data not valid")
	}
	return ou.orderRepository.Create(ctx, orderPrams)
}
