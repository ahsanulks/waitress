package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/ahsanulks/waitress/domain"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase CartUsecase
}

// CartUsecase is all dependency method that needed by cart handler.
type CartUsecase interface {
	FindOrCreate(ctx context.Context, userID int) (domain.Cart, error)
	AddItem(ctx context.Context, cartItemParams domain.CartItemParams) (domain.CartItem, error)
}

// NewCartHandler to create new cart handler.
func NewCartHandler(cartUsecase CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase}
}

// Index handle endpoint GET /.
func (ch CartHandler) Index(c *gin.Context) {
	// TODO: need to change to get userID by jwt token
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		render(c, errors.New("user_id must needed"), http.StatusBadRequest)
		return
	}

	// find cart based on userID, but when cart with userID is not found this will create a new cart for that userID
	cart, err := ch.cartUsecase.FindOrCreate(c, userID)
	if err != nil {
		render(c, err, http.StatusUnprocessableEntity)
		return
	}

	render(c, cart, http.StatusOK)
}

// AddToCart handle endpoint POST /items.
func (ch CartHandler) AddToCart(c *gin.Context) {
	var cartItemParams domain.CartItemParams
	if err := c.ShouldBindJSON(&cartItemParams); err != nil {
		render(c, err, http.StatusBadRequest)
		return
	}

	cartItem, err := ch.cartUsecase.AddItem(c, cartItemParams)
	if err != nil {
		render(c, err, http.StatusUnprocessableEntity)
		return
	}
	render(c, cartItem, http.StatusCreated)
}

// Mount all endpoint to router.
func (ch CartHandler) Mount(router *gin.RouterGroup) {
	router.GET("/", ch.Index)
	router.POST("/items", ch.AddToCart)
}
