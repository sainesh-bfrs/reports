package main

import (
	"reports/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DatabaseMiddleware ...
func DatabaseMiddleware(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func main() {
	db, _ := sqlx.Connect("mysql", "local:local@(localhost:3306)/multichannel_db")

	router := routes.Router()

	router.Use(DatabaseMiddleware(db))

	router.Run()
}
