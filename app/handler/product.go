package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ahsanulks/waitress/domain"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase ProductUsecase
}

type ProductUsecase interface {
	Create(ctx context.Context, product *domain.Product) error
	FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error)
}

func NewProductHandler(productUsecase ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

// Create handle endpoint POST /.
func (ph ProductHandler) Create(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		render(c, err, http.StatusBadRequest)
		return
	}

	if err := ph.productUsecase.Create(c, &product); err != nil {
		render(c, err, http.StatusUnprocessableEntity)
		return
	}

	render(c, product, http.StatusCreated)
}

// Show handle endpoint GET /.
func (ph ProductHandler) Index(c *gin.Context) {
	var (
		limit  = 10
		offset = 0
	)

	if limitParam, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = limitParam
	}

	if offsetParam, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = offsetParam
	}

	products, _ := ph.productUsecase.FindAll(c, limit, offset)
	render(c, products, http.StatusOK)
}

func (ph ProductHandler) Mount(router *gin.RouterGroup) {
	router.POST("/", ph.Create)
	router.GET("/", ph.Index)
}
