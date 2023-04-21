package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
		})

	r.Run()
}
