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

type fakeOrderUsecase struct{}

func (fou fakeOrderUsecase) Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error) {
	if orderPrams.BuyerID != 2 {
		return domain.Order{}, errors.New("error when create")
	}
	return domain.Order{
		ID:      1,
		State:   domain.Pending,
		BuyerID: 2,
	}, nil
}

func TestOrderHandler_Create(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		response string
		status   int
		payload  string
	}{
		{
			name:     "success create",
			path:     "/",
			status:   http.StatusCreated,
			response: `{"data":{"id":1,"code":"","buyer_id":2,"seller_id":0,"state":"pending","total_price":0,"note":"","items":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"meta":{"http_status":201}}`,
			payload:  `{"buyer_id": 2, "cart_item_ids": [1,2]}`,
		},
		{
			name:     "error when create order",
			path:     "/",
			status:   http.StatusUnprocessableEntity,
			response: `{"error":"error when create","meta":{"http_status":422}}`,
			payload:  `{"buyer_id": 3, "cart_item_ids": [1,2]}`,
		},
		{
			name:     "cant parse json",
			path:     "/",
			status:   http.StatusBadRequest,
			response: `{"error":"json: cannot unmarshal string into Go value of type domain.OrderParams","meta":{"http_status":400}}`,
			payload:  `"buyer_id": 2, "cart_item_ids": [1,2]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				router  = gin.New()
				payload = strings.NewReader(tt.payload)
				req, _  = http.NewRequest("POST", tt.path, payload)
				rr      = httptest.NewRecorder()
				handler = handler.NewOrderHandler(fakeOrderUsecase{})
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
