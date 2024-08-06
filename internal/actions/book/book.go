package book

import (
	"log"

	"github.com/gin-gonic/gin"
)

type BookService struct{}

func (b *BookService) RegisterRoutes(r *gin.Engine) {
	r.GET("/books", getAllBooks)
}

// Book Service godoc
// @Summary Retrieve entire books inventory
// @Schemes
// @Description Endpoint to retrieve all books
// @Tags Book
// @Produce json
// @Success 200 {object} []models.Book
// @Router /books [get]
func getAllBooks(c *gin.Context) {
	log.Println("TODO - show entire book inventory")
}
