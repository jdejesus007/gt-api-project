package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdejesus007/gt-api-project/api/provider"
	"github.com/jdejesus007/gt-api-project/internal/constants"
	"github.com/jdejesus007/gt-api-project/internal/models"
)

type customerURIPayload struct {
	CustomerUUID string `json:"customerUUID,omitempty" uri:"customerUUID" binding:"required,uuid"`
}

type orderHistoryResp struct {
	Orders    []*models.Order `json:"orders"`
	Books     []*models.Book  `json:"books"`
	Email     string          `json:"email"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
}

// Order Service godoc
// @Summary Retrieve all orders / order history for particular customer
// @Schemes
// @Description Endpoint to retrieve customer record
// @Tags Customer
// @Produce json
// @Success 200 {object} []models.Order
// @Failure 422
// @Router /customers/{customerUUID}/orders [get]
func customerOrderIndex(c *gin.Context) {
	payload := new(customerURIPayload)
	if err := c.ShouldBindUri(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get orders for customer - err: " + err.Error(),
		})
		return
	}

	var repositories provider.RepositoryProvider
	if repositoriesData, exists := c.Get(constants.DependencyProviderContextKey); exists {
		repositories = repositoriesData.(provider.RepositoryProvider)
	}

	// Ensure customer is valid via email
	var customer *models.Customer
	if result := repositories.Database().GetConn().Table("customers").
		Where("uuid = ?", payload.CustomerUUID).
		Find(&customer); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "failed to get orders for customer - customer not found - err: " + result.Error.Error(),
		})
		return
	}

	var orders []*models.Order
	if result := repositories.Database().GetConn().Table("orders").
		Where("customer_uuid = ?", customer.UUID).
		Find(&orders); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "failed to get orders for customer - sql err: " + result.Error.Error(),
		})
		return
	}

	if len(orders) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	var uuids []string
	for _, o := range orders {
		for _, id := range o.BookUUIDs {
			uuids = append(uuids, id)
		}
	}

	var books []*models.Book
	if result := repositories.Database().GetConn().Table("books").
		Where("uuid IN (?)", uuids).
		Find(&books); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "failed to get orders for customer - sql err: " + result.Error.Error(),
		})
		return
	}

	if len(books) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"error": "failed to get orders for customer - unable to find books for order history",
		})
	}

	c.JSON(http.StatusOK, orderHistoryResp{
		Books:     books,
		Orders:    orders,
		Email:     customer.Email,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
	})
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
	log.Println("TODO - optional - retrieve sole order details for this particular customer")
}
