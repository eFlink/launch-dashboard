package main

import (
	"time"

	"github.com/eFlink/launch-dashboard-server/controller"
	"github.com/eFlink/launch-dashboard-server/initializer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Launch Endpoints
	r.GET("/api/launches", controller.GetAll)
	r.GET("/api/launches/:id", controller.GetByID)
	r.POST("/api/launches", controller.CreateLaunch)
	r.DELETE("/api/launches/:id", controller.DeleteLaunchByID)

	r.Run()
}
