package controllers

import (
	"apitest/db"
	"apitest/models"
	"apitest/services"
	"apitest/utils"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*** เพื่อให้ routes ต่อไปนี้ทำงานต้องนำ บันทึก routes ไปเพิ่มใน server.go
// controllers.SetupBookRoutes(v1)
func SetupBookRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/books")
	// router.Use(utils.JwtVerify) // jwt verify with secret key ถ้าต้องการให้ auth ด้วย jwt
	router.GET("/all", listBook)
	router.GET("/last", listlastBook)
	router.GET("/vuetable", vuetableBook)
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
	payload := utils.GetResponse()
	err := services.GetAllBook(&rs)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
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
	payload := utils.GetResponse()
	err := services.GetAllIdDescBook(&rs)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
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
	payload := utils.GetResponse()
	c.BindJSON(&rs)
	err := services.AddNewBook(&rs)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
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
	payload := utils.GetResponse()
	var rs models.Book
	err := services.GetOneBook(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
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
	payload := utils.GetResponse()
	id := c.Params.ByName("id")
	err := services.GetOneBook(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 200, payload)
	}
	c.BindJSON(&rs)
	err = services.PutOneBook(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 200, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
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
	payload := utils.GetResponse()
	id := c.Params.ByName("id")
	err := services.DeleteBook(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// VuetableBook godoc
// @summary List Books use VueTable
// @description List all books with pagination keyword filter by column Sort columns
// @tags books
// @security ApiKeyAuth
// @id vuetableBook
// @accept json
// @produce json
// @Param page query string false "number of page"
// @Param per_page query string false "get number of per_page"
// @Param sort query string false "soft by cols exm: col1|asc,col2|desc"
// @Param filter query string false "search with column exm: col1|aaa,col2|bbb  by filter"
// @Param kw query string false "search by kw exm: aaa bbb ccc "
// @response 200 {object} []models.Book "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// example /api/v1/books/vuetable?page=2&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=name|aaa,author|bbb
// @Router /api/v1/books/vuetable [get]
func vuetableBook(c *gin.Context) {
	var rs []models.Book
	query := c.Request.URL.Query()
	table := "books"
	paging := utils.GenPagination(query, table)
	dbqry := db.GetDB().Limit(paging.Limit).Offset(paging.Offset).Order(paging.Sort)

	qryfilter := utils.Filter(paging.Filter)
	for k, v := range qryfilter {
		log.Println("K->", k, "V->", v)
		dbqry = dbqry.Where(k+" like ? ", "%"+v+"%")
	}

	// ############################################################ start
	// columns ที่ต้องการให้ค้นหาด้วย keyword รวม
	// cols := []string{"name", "author", "category"} //ต้องกำหนดเอง example
	cols := []string{} //ต้องกำหนดเอง
	// keywork single search #1
	if paging.Kw != "" {
		for _, col := range cols {
			dbqry = dbqry.Or(col+" like ? ", "%"+paging.Kw+"%")
		}
	}
	// keywork multiple search #2
	// var kw = "aaa bbb cccc"
	rskws := utils.Searchkw(paging.Kw)
	log.Println("test kws-->", rskws)
	for _, kw := range rskws {
		for _, col := range cols {
			dbqry = dbqry.Or(col+" like ? ", "%"+kw+"%")
		}
	}
	// ############################################################ end

	dbqry.Find(&rs)
	if dbqry.Error != nil {
		log.Fatal(dbqry.Error)
	}
	var totalRows int
	errCount := db.GetDB().Model(&rs).Count(&totalRows).Error
	if errCount != nil {
		log.Fatal(errCount.Error)
	}

	payload := utils.GetResponse()
	var vtb utils.VueTableResponse
	vtb.Total = totalRows
	vtb.PerPage = paging.PerPage
	vtb.Datas = rs
	vtb.Form = paging.Offset + 1
	vtb.To = paging.Page * paging.Limit
	vtb.CurrentPage = paging.Page
	vtb.LastPage = int(math.Ceil(float64(totalRows) / float64(paging.PerPage)))
	// vtb.NextPageURL = "?page=3&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=field|aaaa,field2|bbbbbb"
	// vtb.PrevPageURL = "?page=1&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=field|aaaa,field2|bbbbbb"
	// vtb := utils.VueTableResponse{Datas: rs, Total: 2000, PerPage: 20, Form: 200, To: 300, CurrentPage: 1, LastPage: 200, NextPageURL: "", PrevPageURL: ""}
	alert := utils.Alert{Msg: "Vue Table สำเร็จ", Title: "Success", Type: "success"}
	payload.Alert = alert
	payload.Data = vtb
	payload.Pagination = paging
	payload.Meta = http.StatusText(200)
	utils.RespondJSON(c, 200, payload)
}
