package customer

import (
	"log"

	"github.com/gin-gonic/gin"
)

type CustomerService struct{}

func (c *CustomerService) RegisterRoutes(r *gin.Engine) {
	cGroup := r.Group("customers")

	cGroup.POST("/", customerCreate)
	cGroup.GET("/:customerUUID", customerShow)
}

// Customer Service godoc
// @Summary Show sole customer by id or email
// @Schemes
// @Description Endpoint to retrieve customer record
// @Tags Customer
// @Produce json
// @Success 200 {object} models.Customer
// @Router /customers/{customerUUID} [get]
func customerShow(c *gin.Context) {

}

// Customer Service godoc
// @Summary Create customer
// @Schemes
// @Description Endpoint to post customer creation
// @Tags Customer
// @Produce json
// @Success 201 {object} models.Customer
// @Router /customers [post]
func customerCreate(c *gin.Context) {
	log.Println("TODO - implement me")
}
