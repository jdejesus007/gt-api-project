package checkout

import (
	"log"

	"github.com/gin-gonic/gin"
)

type CheckoutService struct{}

func (c *CheckoutService) RegisterRoutes(r *gin.Engine) {
	log.Println("TODO - implement checkout and create order within db")
}
