package order

import (
	"log"

	"github.com/gin-gonic/gin"
)

type OrderService struct{}

func (o *OrderService) RegisterRoutes(r *gin.Engine) {
	orderGroup := r.Group("/customers/:customerUUID/orders")

	orderGroup.GET("/", customerOrderIndex)
	orderGroup.GET("/:orderUUID", customerOrderDetails)
}

// Order Service godoc
// @Summary Retrieve all orders for particular customer
// @Schemes
// @Description Endpoint to retrieve customer record
// @Tags Customer
// @Produce json
// @Success 200 {object} []models.Order
// @Router /customers/{customerUUID}/orders [get]
func customerOrderIndex(c *gin.Context) {
	log.Println("TODO - retrieve all orders for this particular customer")
}

// Order Service godoc
// @Summary Retrieve order details
// @Schemes
// @Description Endpoint to retrieve order record
// @Tags Order
// @Produce json
// @Success 200 {object} models.Order
// @Router /customers/{customerUUID}/orders/{orderUUID} [get]
func customerOrderDetails(c *gin.Context) {
	log.Println("TODO - retrieve sole order details for this particular customer")
}
