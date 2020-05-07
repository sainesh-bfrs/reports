package controllers

import (
	"log"
	"reports/utilities"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//UsersController ...
type UsersController struct {
}

// Get ...
func (u *UsersController) Get() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		db, _ := ctx.MustGet("db").(*sqlx.DB)

		rows, err := db.Queryx("select * from users")

		if err != nil {
			log.Print(err)
		}

		a := utilities.MapScanExtended(rows)

		ctx.JSON(200, gin.H{
			"users": a,
		})
	}
	return gin.HandlerFunc(fn)
}
