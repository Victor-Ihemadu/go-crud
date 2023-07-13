package controllers

import (
	"github.com/Victor-Ihemadu/go-crud/initializers"
	"github.com/Victor-Ihemadu/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate (c *gin.Context) {

	//GET DATA OF REQUEST BODY
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//CREATE A POST
	post := models.Post{Title: body.Title, Body: body.Body,}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//RETURN IT.
		c.JSON(200, gin.H{
			"post": post,
		})
}

func PostsIndex(c *gin.Context) {
	//Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
			"posts": posts,
		})
}