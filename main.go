package main

/*
 * File: main.go
 * File Created: Monday, 11th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"reports/config"

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

	// initReceiver()

	// router := routes.Router()

	// router.Run()

}
