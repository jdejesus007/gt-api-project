package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingService struct{}

func (p *PingService) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", pingStatus)
}

// Ping Service godoc
// @Summary Basic health check
// @Schemes
// @Description Endpoint to check for a basic response
// @Tags Ping
// @Produce plain
// @Success 200 {string} pong
// @Router /ping [get]
func pingStatus(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
