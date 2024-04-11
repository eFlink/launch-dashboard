package main

import (
	"github.com/eFlink/launch-dashboard-server/controller"
	"github.com/eFlink/launch-dashboard-server/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Launch Endpoints
	r.GET("/api/launches", controller.GetAll)
	r.GET("/api/launches/:id", controller.GetByID)
	r.POST("/api/launches", controller.CreateLaunch)
	r.DELETE("/api/launches/:id", controller.DeleteLaunchByID)

	r.Run()
}
