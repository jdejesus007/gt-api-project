package customer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdejesus007/gt-api-project/api/provider"
	"github.com/jdejesus007/gt-api-project/internal/constants"
	"github.com/jdejesus007/gt-api-project/internal/models"

	"github.com/google/uuid"
)

type customerPayload struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

// Customer Service godoc
// @Summary Create customer
// @Schemes
// @Description Endpoint to post customer creation
// @Tags Customer
// @Produce json
// @Param firstName body string true "First Name"
// @Param lastName body string true "Last Name"
// @Param email body string true "Email"
// @Success 201 {object} models.Customer
// @Failure 422
// @Failure 400
// @Router /customers [post]
func customerCreate(c *gin.Context) {
	payload := new(customerPayload)
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

	// No need to verify email address per functional requirements
	// Shall error out with duplicate key constraint - enough to rely on this error returning 422
	// Nice to have
	// Look up if existing user exist and gracefully handle the error but we can rely on db constraint
	customer := models.Customer{
		UUID:      uuid.New().String(),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	result := repositories.Database().GetConn().Create(&customer)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, customer)
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
	log.Println("TODO - show sole customer")
}
