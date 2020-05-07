package controllers

import (
	"github.com/gin-gonic/gin"
)

//UsersController ...
type UsersController struct {
}

// Get ...
func (u *UsersController) Get() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		// conn, ok := ctx.MustGet("db").(sqlx.DB)

		// if !ok {
		// 	panic("Error")
		// }

		// rows, err := conn.Queryx("select * from users")

		// if err != nil {
		// 	panic("Error")
		// }

		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	}
	return gin.HandlerFunc(fn)
}
