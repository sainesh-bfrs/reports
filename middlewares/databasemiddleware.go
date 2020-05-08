package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// DatabaseMiddleware ...
func DatabaseMiddleware(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}
