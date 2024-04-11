package main

import (
	"github.com/eFlink/launch-dashboard-server/initializer"
	"github.com/eFlink/launch-dashboard-server/pkg/model"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&model.Launch{})
}
