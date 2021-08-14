package handler

import (
	"github.com/gin-gonic/gin"
)

// Healthz handler.
type Healthz struct{}

// Index handle endpoint GET /.
func (h Healthz) Index(c *gin.Context) {
	render(c, "ok", 200)
}

// Mount healhtz handler to route group.
func (h Healthz) Mount(router *gin.RouterGroup) {
	router.GET("/", h.Index)
}

// NewHealthz create new healthz handler.
func NewHealthz() *Healthz {
	return &Healthz{}
}
