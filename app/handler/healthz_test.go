package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsanulks/waitress/app/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type fakeDependency struct {
	err error
}

func (p fakeDependency) Ping(ctx context.Context) error {
	return p.err
}

func TestHealthz_Index(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		path       string
		response   string
		dependency handler.Checker
	}{
		{
			name:       "everything is ok",
			status:     http.StatusOK,
			path:       "/",
			response:   `{"message":"ok","meta":{"http_status":200}}`,
			dependency: fakeDependency{},
		},
		{
			name:       "have a dependencies error",
			status:     http.StatusServiceUnavailable,
			path:       "/",
			response:   `{"message":"service test_dependency are down","meta":{"http_status":503}}`,
			dependency: fakeDependency{errors.New("error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				router  = gin.New()
				handler = handler.NewHealthz()
				req, _  = http.NewRequest("GET", tt.path, nil)
				rr      = httptest.NewRecorder()
			)

			handler.AddCheck("test_dependency", tt.dependency)
			handler.Mount(router.Group("/"))
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.status, rr.Code)
			assert.JSONEq(t, tt.response, rr.Body.String())
		})
	}
}
