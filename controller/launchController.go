package controller

import (
	"github.com/eFlink/launch-dashboard-server/initializer"
	"github.com/eFlink/launch-dashboard-server/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	var launches []model.Launch
	result := initializer.DB.Find(&launches)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"launches": launches,
	})
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	var launch model.Launch
	initializer.DB.First(&launch, "id = '"+id+"'")
	c.JSON(200, gin.H{
		"id": launch,
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
		"launch": launch,
	})
}

func DeleteLaunchByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Delete a launch",
		"id":      id,
	})
}
