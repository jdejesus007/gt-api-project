package internal

import (
	"github.com/gin-gonic/gin"
)

type FavIconService struct{}

func (f *FavIconService) RegisterRoutes(r *gin.Engine) {
	r.StaticFile("favicon.ico", "./assets/gopher.png")
}
