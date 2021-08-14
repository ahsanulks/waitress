package handler

import (
	"context"
	"net/http"

	"github.com/ahsanulks/waitress/domain"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productWriter ProductWriter
}

type ProductWriter interface {
	Create(ctx context.Context, product *domain.Product) error
}

func NewProductHandler(productWriter ProductWriter) *ProductHandler {
	return &ProductHandler{productWriter: productWriter}
}

func (ph ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		render(c, err, http.StatusBadRequest)
		return
	}

	if err := ph.productWriter.Create(c, &product); err != nil {
		render(c, err, http.StatusUnprocessableEntity)
		return
	}

	render(c, product, http.StatusCreated)
}

func (ph ProductHandler) Mount(router *gin.RouterGroup) {
	router.POST("/", ph.CreateProduct)
}
