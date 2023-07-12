package main

import (
	"github.com/Victor-Ihemadu/go-crud/initializers"
	"github.com/Victor-Ihemadu/go-crud/models"
)

func init() {
	initializers.LoadENvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}