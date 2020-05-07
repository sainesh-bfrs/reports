package routes

import (
	"reports/controllers"

	"github.com/gin-gonic/gin"
)

//Router ...
func Router() *gin.Engine {
	var u *controllers.UsersController
	var c *controllers.CompaniesController
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/users", u.Get())
		v1.GET("/companies", c.Get())
	}
	return r
}
