package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahsanulks/waitress/app/handler"
	"github.com/ahsanulks/waitress/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type fakeCartUsecase struct{}

func (fcu fakeCartUsecase) FindOrCreate(ctx context.Context, userID int) (domain.Cart, error) {
	if userID == 2 {
		return domain.Cart{}, errors.New("something happen")
	}
	return domain.Cart{ID: 1}, nil
}

func (fcu fakeCartUsecase) AddItem(ctx context.Context, cartItemParams domain.CartItemParams) (domain.CartItem, error) {
	if cartItemParams.CartID == 10 {
		return domain.CartItem{}, errors.New("error add item")
	}
	return domain.CartItem{ID: 10}, nil
}

func TestCartHandler_Index(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		response string
		status   int
	}{
		{
			name:     "request without user_id",
			path:     "/",
			response: `{"error":"user_id must needed","meta":{"http_status":400}}`,
			status:   http.StatusBadRequest,
		},
		{
			name:     "error when find or create",
			path:     "/?user_id=2",
			response: `{"error":"something happen","meta":{"http_status":422}}`,
			status:   http.StatusUnprocessableEntity,
		},
		{
			name:     "success fully create",
			path:     "/?user_id=1",
			response: `{"data":{"created_at":"0001-01-01T00:00:00Z", "id":1, "updated_at":"0001-01-01T00:00:00Z", "user_id":0}, "meta":{"http_status":200}}`,
			status:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				router  = gin.New()
				req, _  = http.NewRequest("GET", tt.path, nil)
				rr      = httptest.NewRecorder()
				handler = handler.NewCartHandler(fakeCartUsecase{})
			)

			handler.Mount(router.Group("/"))
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.status, rr.Code)
			if tt.response != "" {
				assert.JSONEq(t, tt.response, rr.Body.String())
			} else {
				assert.Equal(t, "", rr.Body.String())
			}
		})
	}
}

func TestCartHandler_AddToCart(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		response string
		status   int
		payload  string
	}{
		{
			name:     "error when binding json",
			path:     "/items",
			response: `{"error":"json: cannot unmarshal string into Go value of type domain.CartItemParams","meta":{"http_status":400}}`,
			status:   http.StatusBadRequest,
			payload:  `"cart_id": 1`,
		},
		{
			name:     "error when adding item",
			path:     "/items",
			response: `{"error":"error add item","meta":{"http_status":422}}`,
			status:   http.StatusUnprocessableEntity,
			payload:  `{"cart_id": 10}`,
		},
		{
			name:     "success create item",
			path:     "/items",
			response: `{"data":{"created_at":"0001-01-01T00:00:00Z", "id":10, "purchased":false, "quantity":0, "updated_at":"0001-01-01T00:00:00Z"}, "meta":{"http_status":201}}`,
			status:   http.StatusCreated,
			payload:  `{"cart_id": 2}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				router  = gin.New()
				payload = strings.NewReader(tt.payload)
				req, _  = http.NewRequest("POST", tt.path, payload)
				rr      = httptest.NewRecorder()
				handler = handler.NewCartHandler(fakeCartUsecase{})
			)

			handler.Mount(router.Group("/"))
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.status, rr.Code)
			if tt.response != "" {
				assert.JSONEq(t, tt.response, rr.Body.String())
			} else {
				assert.Equal(t, "", rr.Body.String())
			}
		})
	}
}
