package checkout

import "github.com/gin-gonic/gin"

type CheckoutService struct{}

func (c *CheckoutService) RegisterRoutes(r *gin.Engine) {
	r.POST("/checkout", checkout)
}
