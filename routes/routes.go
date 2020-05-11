package routes

import (
	"reports/controllers"

	"github.com/gin-gonic/gin"
)

var bookCtl = new(controllers.BookController)

// Router ...
func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	v1 := r.Group("v1")
	{
		v1.GET("/book", bookCtl.ListBook)
		v1.POST("/book", bookCtl.AddNewBook)
		v1.GET("/book/:id", bookCtl.GetOneBook)
		v1.PUT("/book/:id", bookCtl.PutOneBook)
		v1.DELETE("/book/:id", bookCtl.DeleteBook)
	}
	return r
}
