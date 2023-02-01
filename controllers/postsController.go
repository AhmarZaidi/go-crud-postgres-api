package controllers

import (
	"github.com/AhmarZaidi/go-crud-postgres/initializers"
	"github.com/AhmarZaidi/go-crud-postgres/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data of req body

 
	// Create a post
	post := models.Post{Title: "Hello (Title)", Body: "World (Body)"}
	result := initializers.DB.Create(&post) // pass pointer of data to Create
	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	// Error
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}