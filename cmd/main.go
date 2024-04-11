package main

import (
	"github.com/eFlink/launch-dashboard-server/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/api/launches", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List Launches",
		})
	})

	r.GET("/api/launches/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "List Launch",
			"id":      id,
		})
	})

	r.POST("/api/launches", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Post a launch",
		})
	})

	r.DELETE("/api/launches/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Delete a launch",
			"id":      id,
		})
	})

	r.Run()
}
