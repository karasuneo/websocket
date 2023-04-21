package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karasuneo/websocket/backend/server/handlers"
)

func Init() {
	r := gin.Default()

	r.GET("/",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
		})

	r.GET("/socket", handlers.NewWebsocketHandler().Handle)

	r.Run()
}
