package book

import "github.com/gin-gonic/gin"

type BookService struct{}

func (b *BookService) RegisterRoutes(r *gin.Engine) {
	r.GET("/books", getAllBooks)
}
