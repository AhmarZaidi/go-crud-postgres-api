package main

import (
	"github.com/AhmarZaidi/go-crud-postgres/controllers"
	"github.com/AhmarZaidi/go-crud-postgres/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	// CREATE
	r.POST("/posts", controllers.PostsCreate)

	// READ
	r.GET("/posts", controllers.PostsReadAll)
	r.GET("/posts/:id", controllers.PostsRead)

	// UPDATE
	r.PUT("/posts/:id", controllers.PostsUpdate)

	// DELETE
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
