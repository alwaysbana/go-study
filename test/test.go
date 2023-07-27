package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		name := c.Query("query")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}
