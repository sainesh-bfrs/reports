package main

import (
	"log"
	"reports/conf"
	"reports/middlewares"
	"reports/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")

	host := conf.Get("DB_HOST")
	user := conf.Get("DB_USER")
	password := conf.Get("DB_PASSWORD")
	dbname := conf.Get("DB_NAME")
	port := conf.Get("DB_PORT")

	db, err := sqlx.Connect("mysql", user+":"+password+"@("+host+":"+port+")/"+dbname)

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.Use(middlewares.DatabaseMiddleware(db))

	router := routes.Router(r)

	router.Run()
}
