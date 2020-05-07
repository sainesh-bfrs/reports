package routes

import (
	"reports/controllers"

	"github.com/gin-gonic/gin"
)

//Router ...
func Router(r *gin.Engine) *gin.Engine {
	var u *controllers.UsersController
	var c *controllers.CompaniesController
	v1 := r.Group("v1")
	{
		v1.GET("/users", u.Get())     // /v1/users
		v1.GET("/companies", c.Get()) // /v1/companies
	}
	return r
}
