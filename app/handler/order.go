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

type OrderUsecase interface {
	Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error)
}

func NewOrderHandler(orderUsecase OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

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

func (oh OrderHandler) Mount(router *gin.RouterGroup) {
	router.POST("/", oh.Create)
}
