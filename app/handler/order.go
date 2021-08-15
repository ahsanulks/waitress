package handler

import (
	"context"
	"net/http"

	"github.com/ahsanulks/waitress/domain"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase OrderUsecase
}

// OrderUsecase method that needed for order handler.
type OrderUsecase interface {
	Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error)
}

// NewOrderHandler will create new order handler
func NewOrderHandler(orderUsecase OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

// Create handle endpoint POST /.
func (oh OrderHandler) Create(c *gin.Context) {
	var orderPrams domain.OrderParams
	if err := c.ShouldBindJSON(&orderPrams); err != nil {
		render(c, err, http.StatusBadRequest)
		return
	}

	order, err := oh.orderUsecase.Create(c, orderPrams)
	if err != nil {
		render(c, err, http.StatusUnprocessableEntity)
		return
	}
	render(c, order, http.StatusCreated)
}

// Mount all endpoint to router
func (oh OrderHandler) Mount(router *gin.RouterGroup) {
	router.POST("/", oh.Create)
}
