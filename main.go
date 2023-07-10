package main

import (
	"github.com/Victor-Ihemadu/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}

