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
// controllers.SetupPostRoutes(v1)
func SetupPostRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/posts")
	// router.Use(utils.JwtVerify) // jwt verify with secret key ถ้าต้องการให้ auth ด้วย jwt
	router.GET("/aaa", aaaa)
	router.GET("/all", listPost)
	router.GET("/last", listlastPost)
	router.GET("/vuetable", vuetablePost)
	router.GET("/by/:id", getOnePost)
	router.GET("/del/:id", deletePost)
	router.POST("/new", addNewPost)
	router.POST("/edit/:id", putOnePost)
}

// ListPosts godoc
// @summary List Posts
// @description List all posts
// @tags posts
// @security ApiKeyAuth
// @id Aaaa
// @accept json
// @produce json
// @response 200 {array} []models.Post "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/aaa [get]
func aaaa(c *gin.Context) {
	payload := utils.GetResponse()
	payload.Data = "Test  A"
	utils.RespondJSON(c, 200, payload)
}

// ListPosts godoc
// @summary List Posts
// @description List all posts
// @tags posts
// @security ApiKeyAuth
// @id ListPosts
// @accept json
// @produce json
// @response 200 {array} []models.Post "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/all [get]
func listPost(c *gin.Context) {
	var rs []models.Post
	payload := utils.GetResponse()
	err := services.GetAllPost(&rs)
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

// ListPosts godoc
// @summary List Posts
// @description List all posts
// @tags posts
// @security ApiKeyAuth
// @id ListlastPosts
// @accept json
// @produce json
// @response 200 {array} []models.Post "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/last [get]
func listlastPost(c *gin.Context) {
	var rs []models.Post
	payload := utils.GetResponse()
	err := services.GetAllIdDescPost(&rs)
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

// CreatePost godoc
// @summary Create Post
// @description Create new post
// @tags posts
// @security ApiKeyAuth
// @id CreatePost
// @accept json
// @produce json
// @param Post body models.PostForCreate true "Post data to be created"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/new [post]
func addNewPost(c *gin.Context) {
	var rs models.Post
	payload := utils.GetResponse()
	c.BindJSON(&rs)
	err := services.AddNewPost(&rs)
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

// GetPost godoc
// @summary Get Post
// @description  Get post by id
// @tags posts
// @security Basic auth
// @id GetPost
// @accept json
// @produce json
// @param id path int true "id of post to be gotten"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/by/{id} [get]
func getOnePost(c *gin.Context) {
	id := c.Params.ByName("id")
	payload := utils.GetResponse()
	var rs models.Post
	err := services.GetOnePost(&rs, id)
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

// UpdatePost godoc
// @summary Update Post
// @description Update post by id
// @tags posts
// @security ApiKeyAuth
// @id UpdatePost
// @accept json
// @produce json
// @param id path int true "id of post to be updated"
// @param Post body models.PostForUpdate true "Post data to be updated"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/edit/{id} [post]
func putOnePost(c *gin.Context) {
	var rs models.Post
	payload := utils.GetResponse()
	id := c.Params.ByName("id")
	err := services.GetOnePost(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 200, payload)
	}
	c.BindJSON(&rs)
	err = services.PutOnePost(&rs, id)
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

// DeletePost godoc
// @summary Delete Post
// @description Delete post by id
// @tags posts
// @security ApiKeyAuth
// @id DeletePost
// @accept json
// @produce json
// @param id path int true "id of post to be deleted"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/posts/del/{id} [get]
func deletePost(c *gin.Context) {
	var rs models.Post
	payload := utils.GetResponse()
	id := c.Params.ByName("id")
	err := services.DeletePost(&rs, id)
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

// VuetablePost godoc
// @summary List Posts use VueTable
// @description List all posts with pagination keyword filter by column Sort columns
// @tags posts
// @security ApiKeyAuth
// @id vuetablePost
// @accept json
// @produce json
// @Param page query string false "number of page"
// @Param per_page query string false "get number of per_page"
// @Param sort query string false "soft by cols exm: col1|asc,col2|desc"
// @Param filter query string false "search with column exm: col1|aaa,col2|bbb  by filter"
// @Param kw query string false "search by kw exm: aaa bbb ccc "
// @response 200 {object} []models.Post "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// example /api/v1/posts/vuetable?page=2&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=name|aaa,author|bbb
// @Router /api/v1/posts/vuetable [get]
func vuetablePost(c *gin.Context) {
	var rs []models.Post
	query := c.Request.URL.Query()
	table := "posts"
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
