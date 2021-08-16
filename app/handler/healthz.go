package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Checker interface for handle all dependencies have a ping method to check that dependencies is healthy.
type Checker interface {
	Ping(ctx context.Context) error
}

// Healthz handler.
type HealthzHandler struct {
	checkers map[string]Checker
}

// Index handle endpoint GET /.
func (h HealthzHandler) Index(c *gin.Context) {
	var (
		status  = http.StatusOK
		message = "ok"
	)

	for service, checker := range h.checkers {
		if err := checker.Ping(c); err != nil {
			status = http.StatusServiceUnavailable
			message = fmt.Sprintf("service %s are down", service)
			break
		}
	}

	render(c, message, status)
}

// Mount healhtz handler to route group.
func (h HealthzHandler) Mount(router *gin.RouterGroup) {
	router.GET("/", h.Index)
}

// AddCheck is to add dependencies that whant to check on healthz.
func (h *HealthzHandler) AddCheck(serviceName string, service Checker) {
	h.checkers[serviceName] = service
}

// NewHealthz create new healthz handler.
func NewHealthz() *HealthzHandler {
	return &HealthzHandler{
		checkers: make(map[string]Checker),
	}
}
