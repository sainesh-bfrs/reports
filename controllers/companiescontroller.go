package controllers

import (
	"github.com/gin-gonic/gin"
)

// CompaniesController ...
type CompaniesController struct {
}

// Get ...
func (c *CompaniesController) Get() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		// Do Some Stuff
	}
	return gin.HandlerFunc(fn)
}
