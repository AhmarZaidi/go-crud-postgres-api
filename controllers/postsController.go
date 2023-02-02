package controllers

import (
	"github.com/AhmarZaidi/go-crud-postgres/initializers"
	"github.com/AhmarZaidi/go-crud-postgres/models"
	"github.com/gin-gonic/gin"
)

// CREATE
func PostsCreate(c *gin.Context) {

	// Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
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

// READ
func PostsReadAll(c *gin.Context) {

	var posts []models.Post
	// Get all records
	initializers.DB.Find(&posts)
	// SELECT * FROM users;
	// result.RowsAffected // returns found records count, equals `len(users)`
	// result.Error        // returns error

	// Return post
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsRead(c *gin.Context) {

	id := c.Param("id")

	// Find the post to update
	var post models.Post
	result := initializers.DB.First(&post, id) // SELECT * FROM users WHERE id = 10;

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

// Update
func PostsUpdate(c *gin.Context) {
	// Get id from the url
	id := c.Param("id")

	// Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post to update
	var post models.Post
	result := initializers.DB.First(&post, id) // SELECT * FROM users WHERE id = 10;

	// Error
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Update
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	}) // UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// DELETE
func PostsDelete(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Delete Post
	result := initializers.DB.Delete(&models.Post{}, id) // DELETE FROM users WHERE id = 10;

	// Error
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond
	c.Status(200)
}
