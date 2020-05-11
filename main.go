package main

import (
	"reports/config"
	"reports/routes"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	// load configurations
	serverConfig := config.Config

	//  initialize all configurations
	// config.InitDB(serverConfig)
	// defer config.DB.Close()

	gin.SetMode(serverConfig.RunMode)
	//gin.DisableConsoleColor()

	router := routes.Router()

	router.Run()

}
