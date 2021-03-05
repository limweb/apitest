package controllers

import (
	"apitest/model"
	"apitest/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(rg *gin.RouterGroup) {
	books := rg.Group("/books")
	books.GET("", listBook)
	books.POST("", addNewBook)
	books.GET("/:id", getOneBook)
	books.PUT("", putOneBook)
	books.DELETE("/:id", deleteBook)
}

// listBook godoc
// @summary List Book
// @description List all Books
// @tags books
// @security ApiKeyAuth
// @id listBook
// @accept json
// @produce json
// @response 200 {array}  model.Book "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 409 {object} model.Response "Conflict"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/books [get]
func listBook(c *gin.Context) {
	var book []model.Book
	err := services.GetAllBook(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": book,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": book,
		})
	}
}

func addNewBook(c *gin.Context) {
	var book model.Book
	c.BindJSON(&book)
	err := services.AddNewBook(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": book,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": book,
		})
	}
}

func getOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book model.Book
	err := services.GetOneBook(&book, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": book,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": book,
		})
	}
}

func putOneBook(c *gin.Context) {
	var book model.Book
	id := c.Params.ByName("id")
	err := services.GetOneBook(&book, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": book,
		})
	}
	c.BindJSON(&book)
	err = services.PutOneBook(&book, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": book,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": book,
		})
	}
}

func deleteBook(c *gin.Context) {
	var book model.Book
	id := c.Params.ByName("id")
	err := services.DeleteBook(&book, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": book,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": book,
		})
	}
}
