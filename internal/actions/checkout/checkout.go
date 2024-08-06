package checkout

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jdejesus007/gt-api-project/api/provider"
	"github.com/jdejesus007/gt-api-project/internal/constants"
	"github.com/jdejesus007/gt-api-project/internal/models"
)

type checkoutPayload struct {
	Email string         `json:"email" binding:"required"`
	Books []*models.Book `json:"books" binding:"required"`
}

// Checkout Service godoc
// @Summary checkout the customer with list of books
// @Schemes
// @Description Endpoint to check out customer
// @Tags Checkout
// @Produce json
// @Param books body array true "List of Books to Create Order With"
// @Param email body string true "Email"
// @Success 201 {object} models.Order
// @Failure 422
// @Router /checkout [post]
func checkout(c *gin.Context) {
	// Customer account must be created prior to checkout
	// Ensure customer already exists to checkout
	log.Println("TODO - check out customer with list of books")

	payload := new(checkoutPayload)
	if err := c.ShouldBindJSON(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	var repositories provider.RepositoryProvider
	if repositoriesData, exists := c.Get(constants.DependencyProviderContextKey); exists {
		repositories = repositoriesData.(provider.RepositoryProvider)
	}

	var customer *models.Customer
	if result := repositories.Database().GetConn().Table("customers").
		Where("email = ? and base_status <> ?", payload.Email, models.BaseStatusEnumDeleted).
		Find(&customer); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "customer not fount - unable to checkout",
		})
		return
	}

	if customer.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Verify books are still in stock
	var uuids []string
	for _, b := range payload.Books {
		uuids = append(uuids, b.UUID)
	}
	var dbBooks []models.Book
	if result := repositories.Database().GetConn().Table("books").
		Where("uuid IN (?) and base_status <> ?", uuids, models.BaseStatusEnumDeleted).
		Find(&dbBooks); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "failed to verify book invenvory - unable to checkout - error: " + result.Error.Error(),
		})
		return
	}

	// Match up the found inventory to the payload
	// Any inconsistencies found - error out with 422
	// Nicety: error out with list of invalid items
	var (
		invalidItem bool
	)
	for _, b := range payload.Books {
		found := false
		for _, dbB := range dbBooks {
			if b.UUID == dbB.UUID {
				found = true
				break
			}
		}

		// Exit on first invalid item and return 422
		if !found {
			invalidItem = true
			break
		}
	}

	if invalidItem {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "unable to checkout due to invalid books",
		})
	}

	// Able to checkout
	// Create order and return obj
	order := models.Order{
		UUID:         uuid.New().String(),
		CustomerUUID: customer.UUID,
		BookUUIDs:    uuids,
	}

	result := repositories.Database().GetConn().Create(&order)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "unable to checkout - error: " + result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, order)
}
