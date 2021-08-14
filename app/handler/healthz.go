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

type check struct {
	Service string
	Status  string
}

// Healthz handler.
type Healthz struct {
	checkers map[string]Checker
}

// Index handle endpoint GET /.
func (h Healthz) Index(c *gin.Context) {
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
func (h Healthz) Mount(router *gin.RouterGroup) {
	router.GET("/", h.Index)
}

func (h *Healthz) AddCheck(serviceName string, service Checker) {
	h.checkers[serviceName] = service
}

// NewHealthz create new healthz handler.
func NewHealthz() *Healthz {
	return &Healthz{
		checkers: make(map[string]Checker),
	}
}
