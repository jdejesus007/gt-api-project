package checkout

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Checkout Service godoc
// @Summary checkout the customer with list of books
// @Schemes
// @Description Endpoint to check out customer
// @Tags Checkout
// @Produce json
// @Success 201 {object} models.Order
// @Failure 422
// @Router /checkout [post]
func checkout(c *gin.Context) {
	log.Println("TODO - check out customer with list of books")
}
