package order

import "github.com/gin-gonic/gin"

type OrderService struct{}

func (o *OrderService) RegisterRoutes(r *gin.Engine) {
	orderGroup := r.Group("/customers/:customerUUID/orders")

	orderGroup.GET("/", customerOrderIndex)
	orderGroup.GET("/:orderUUID", customerOrderDetails)
}
