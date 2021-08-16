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

type fakeProductUsecase struct{}

func (fpu fakeProductUsecase) Create(ctx context.Context, product *domain.Product) error {
	if product.SellerID == 1 {
		return errors.New("can't create product")
	}
	return nil
}

func (fpu fakeProductUsecase) FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	if limit == 1 && offset == 100 {
		return []domain.Product{}, nil
	} else if limit == 1 && offset == 1 {
		return []domain.Product{
			{
				ID: 2,
			},
		}, nil
	}
	return []domain.Product{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}, nil
}

func TestProductHandler_Create(t *testing.T) {
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
			response: `{"data":{"id":0,"name":"","seller_id":2,"price":0,"active":false,"stock":0,"weight":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"meta":{"http_status":201}}`,
			status:   http.StatusCreated,
			payload:  `{"seller_id": 2}`,
		},
		{
			name:     "error parse json ",
			path:     "/",
			response: `{"error":"json: cannot unmarshal string into Go value of type domain.Product","meta":{"http_status":400}}`,
			status:   http.StatusBadRequest,
			payload:  `"seller_id": 2`,
		},
		{
			name:     "error when create product",
			path:     "/",
			response: `{"error":"can't create product","meta":{"http_status":422}}`,
			status:   http.StatusUnprocessableEntity,
			payload:  `{"seller_id": 1}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				router  = gin.New()
				payload = strings.NewReader(tt.payload)
				req, _  = http.NewRequest("POST", tt.path, payload)
				rr      = httptest.NewRecorder()
				handler = handler.NewProductHandler(fakeProductUsecase{})
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

func TestProductHandler_Index(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		response string
		status   int
	}{
		{
			name:     "access without query",
			path:     "/",
			response: `{"data":[{"id":1,"name":"","seller_id":0,"price":0,"active":false,"stock":0,"weight":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},{"id":2,"name":"","seller_id":0,"price":0,"active":false,"stock":0,"weight":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}],"meta":{"http_status":200}}`,
			status:   http.StatusOK,
		},
		{
			name:     "access with query",
			path:     "/?offset=1&limit=1",
			response: `{"data":[{"id":2,"name":"","seller_id":0,"price":0,"active":false,"stock":0,"weight":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}],"meta":{"http_status":200}}`,
			status:   http.StatusOK,
		},
		{
			name:     "not found",
			path:     "/?offset=100&limit=1",
			response: `{"data":[],"meta":{"http_status":200}}`,
			status:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				router  = gin.New()
				req, _  = http.NewRequest("GET", tt.path, nil)
				rr      = httptest.NewRecorder()
				handler = handler.NewProductHandler(fakeProductUsecase{})
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
