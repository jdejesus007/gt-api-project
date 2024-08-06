package customer

import "github.com/gin-gonic/gin"

type CustomerService struct{}

func (c *CustomerService) RegisterRoutes(r *gin.Engine) {
	cGroup := r.Group("customers")

	cGroup.POST("/", customerCreate)
	cGroup.GET("/:customerUUID", customerShow)
}
