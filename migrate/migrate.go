package main

import (
	"github.com/AhmarZaidi/go-crud-postgres/initializers"
	"github.com/AhmarZaidi/go-crud-postgres/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}