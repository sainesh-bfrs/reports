package controllers

/*
 * @Script: book.go
 * @Author: Jayanta
 * @Description: BookController
 */

import (
	"fmt"
	"reports/helpers"
	"reports/models"

	"github.com/gin-gonic/gin"
)

type BookController struct {
}

func (i *BookController) ListBook(c *gin.Context) {
	var book []models.Book
	err := models.GetAllBook(&book)
	fmt.Println(err)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}

func (i *BookController) AddNewBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.AddNewBook(&book)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}

func (i *BookController) GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	err := models.GetOneBook(&book, id)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}

func (i *BookController) PutOneBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.GetOneBook(&book, id)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = models.PutOneBook(&book, id)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}

func (i *BookController) DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.DeleteBook(&book, id)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}
