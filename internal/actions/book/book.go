package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdejesus007/gt-api-project/api/provider"
	"github.com/jdejesus007/gt-api-project/internal/constants"
	"github.com/jdejesus007/gt-api-project/internal/models"
)

// Book Service godoc
// @Summary Retrieve entire books inventory
// @Schemes
// @Description Endpoint to retrieve all books
// @Tags Book
// @Produce json
// @Success 200 {object} []models.Book
// @Failure 404
// @Router /books [get]
func getAllBooks(c *gin.Context) {
	var repositories provider.RepositoryProvider
	if repositoriesData, exists := c.Get(constants.DependencyProviderContextKey); exists {
		repositories = repositoriesData.(provider.RepositoryProvider)
	}

	var books []*models.Book
	repositories.Database().GetConn().Find(&books)
	if len(books) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"error": "No books",
		})
		return
	}

	c.JSON(http.StatusOK, books)
}
