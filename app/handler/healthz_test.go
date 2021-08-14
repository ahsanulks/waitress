package handler_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/app/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewHealthz(t *testing.T) {
	tests := []struct {
		name string
		want *handler.Healthz
	}{
		{
			name: "should return healthz",
			want: &handler.Healthz{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.NewHealthz(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealthz_Index(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		path     string
		response string
	}{
		{
			name:     "everything is ok",
			status:   http.StatusOK,
			path:     "/",
			response: `{"message":"ok"}`,
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

			handler.Mount(router.Group("/"))
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.status, rr.Code)
			assert.JSONEq(t, tt.response, rr.Body.String())
		})
	}
}
