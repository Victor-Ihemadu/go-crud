package main

import (
	"github.com/Victor-Ihemadu/go-crud/controllers"
	"github.com/Victor-Ihemadu/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	
	r.Run()
}

