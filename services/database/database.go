package database

/*
 * File: database.go
 * File Created: Monday, 11th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"fmt"
	"log"
	"reports/config"
	"reports/helpers"

	"github.com/jinzhu/gorm"

	// Loading MySQL Drivers
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB ...
var DB *gorm.DB

// DbURL ...
func dbURL(cfg *config.ServerConfig) string {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.DbName,
	)
}

// InitDB ...
func init() {

	var err error
	var cfg = config.Config
	DB, err = gorm.Open("mysql", dbURL(cfg))

	helpers.LogError("Error in creating connection", err)

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	log.Println("database init on: ", cfg.Host)
}
