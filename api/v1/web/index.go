package web

import "github.com/gin-gonic/gin"

type WebApi struct{}

func (w *WebApi) IndexAd(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
