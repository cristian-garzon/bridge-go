package main

import (
	"awesomeProject/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/play", func(c *gin.Context) {
		playerType := c.GetHeader("media-player")
		utils.GetPlayerByHeader(playerType)
		c.JSONP(http.StatusOK, gin.H{
			"detail": "ok",
		})
		return
	})

	router.Run(":8080")
}
