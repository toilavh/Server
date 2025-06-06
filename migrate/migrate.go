package main

import (
	"awesome_project/initializers"
	"awesome_project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Robot{})
}
