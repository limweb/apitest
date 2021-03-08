package controllers

import (
	"apitest/models"
	"apitest/services"
	"apitest/utils"

	"github.com/gin-gonic/gin"
)

//*** เพื่อให้ routes ต่อไปนี้ทำงานต้องนำ บันทึก routes ไปเพิ่มใน server.go
// controllers.SetupBookRoutes(v1)
func SetupBookRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/books")
	router.Use(utils.JwtVerify) // jwt verify with secret key ถ้าต้องการให้ auth ด้วย jwt
	router.GET("/all", listBook)
	router.GET("/last", listlastBook)
	router.GET("/by/:id", getOneBook)
	router.GET("/del/:id", deleteBook)
	router.POST("/new", addNewBook)
	router.POST("/edit/:id", putOneBook)
}

// ListBooks godoc
// @summary List Books
// @description List all books
// @tags books
// @security ApiKeyAuth
// @id ListBooks
// @accept json
// @produce json
// @response 200 {array} []models.Book "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/books/all [get]
func listBook(c *gin.Context) {
	var rs []models.Book
	err := services.GetAllBook(&rs)
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		utils.RespondJSON(c, 404, []string{}, alert)
	} else {
		utils.RespondJSON(c, 200, rs, alert)
	}
}

// ListBooks godoc
// @summary List Books
// @description List all books
// @tags books
// @security ApiKeyAuth
// @id ListlastBooks
// @accept json
// @produce json
// @response 200 {array} []models.Book "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/books/last [get]
func listlastBook(c *gin.Context) {
	var rs []models.Book
	err := services.GetAllIdDescBook(&rs)
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		utils.RespondJSON(c, 404, []string{}, alert)
	} else {
		utils.RespondJSON(c, 200, rs, alert)
	}
}

// CreateBook godoc
// @summary Create Book
// @description Create new book
// @tags books
// @security ApiKeyAuth
// @id CreateBook
// @accept json
// @produce json
// @param Book body models.BookForCreate true "Book data to be created"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/books/new [post]
func addNewBook(c *gin.Context) {
	var rs models.Book
	c.BindJSON(&rs)
	err := services.AddNewBook(&rs)
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		utils.RespondJSON(c, 404, []string{}, alert)
	} else {
		utils.RespondJSON(c, 200, rs, alert)
	}
}

// GetBook godoc
// @summary Get Book
// @description  Get book by id
// @tags books
// @security Basic auth
// @id GetBook
// @accept json
// @produce json
// @param id path int true "id of book to be gotten"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/books/by/{id} [get]
func getOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var rs models.Book
	err := services.GetOneBook(&rs, id)
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		utils.RespondJSON(c, 404, []string{}, alert)
	} else {
		utils.RespondJSON(c, 200, rs, alert)
	}
}

// UpdateBook godoc
// @summary Update Book
// @description Update book by id
// @tags books
// @security ApiKeyAuth
// @id UpdateBook
// @accept json
// @produce json
// @param id path int true "id of book to be updated"
// @param Book body models.BookForUpdate true "Book data to be updated"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/books/edit/{id} [post]
func putOneBook(c *gin.Context) {
	var rs models.Book
	id := c.Params.ByName("id")
	err := services.GetOneBook(&rs, id)
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Success", Type: "success"}
		utils.RespondJSON(c, 404, []string{}, alert)
	}
	c.BindJSON(&rs)
	err = services.PutOneBook(&rs, id)
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		utils.RespondJSON(c, 404, []string{}, alert)
	} else {
		utils.RespondJSON(c, 200, rs, alert)
	}
}

// DeleteBook godoc
// @summary Delete Book
// @description Delete book by id
// @tags books
// @security ApiKeyAuth
// @id DeleteBook
// @accept json
// @produce json
// @param id path int true "id of book to be deleted"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/books/del/{id} [get]
func deleteBook(c *gin.Context) {
	var rs models.Book
	id := c.Params.ByName("id")
	err := services.DeleteBook(&rs, id)
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	if err != nil {
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		utils.RespondJSON(c, 404, []string{}, alert)
	} else {
		utils.RespondJSON(c, 200, rs, alert)
	}
}
