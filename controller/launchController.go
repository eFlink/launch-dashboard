package controller

import (
	"github.com/eFlink/launch-dashboard-server/initializer"
	"github.com/eFlink/launch-dashboard-server/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	// Get data of req body

	// Get all the launches

	// Return it

	c.JSON(200, gin.H{
		"message": "List Launches",
	})
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "List Launch",
		"id":      id,
	})
}
func CreateLaunch(c *gin.Context) {

	launch := model.Launch{
		Name: "Test 1",
	}

	result := initializer.DB.Create(&launch)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": launch,
	})
}

func DeleteLaunchByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Delete a launch",
		"id":      id,
	})
}
