package database

/*
 * @Script: database.go
 * @Author: Sainesh
 * @Description: database service.
 */

import (
	"fmt"
	"log"
	"reports/config"

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
	DB, _ = gorm.Open("mysql", dbURL(cfg))
	if err != nil {
		log.Fatal(2, err)
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	log.Println("database init on: ", cfg.Host)
}
