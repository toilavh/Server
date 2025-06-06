package main

import (
	"awesome_project/controllers"
	"awesome_project/initializers"
	"awesome_project/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {
	r := gin.Default()

	// Middleware để kiểm tra API key
	r.Use(middlewares.RequireAPIKey())

	// Routes cho người dùng
	r.POST("/robots", controllers.CreateRobot)
	r.GET("/robots", controllers.GetAllRobots)
	r.DELETE("/robots/:id", controllers.DeleteRobot)

	r.Run()

}
