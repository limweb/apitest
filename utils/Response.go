package utils

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// type Response struct {
// 	Code   int    `json:code`   // Response Code
// 	Status string `json:status` // Response Status
// }

type Alert struct {
	Title string `json:"title"  example:"Success" `   // Response title
	Type  string `json:"type"  example:"success" `    // Response type
	Msg   string `json:"msg"  example:"ค้นหาสำเร็จ" ` // Response msg
}

type Pagination struct {
	Sort    string      `json:"sort"  example:"col1|ase,col2|desc,col3|asc" ` // pagination sort
	Kw      string      `json:"kw" example:"kw1 kw2 kw2" `                    // pagination kw
	Filter  string      `json:"filter"`                                       // pagination filter
	Page    int         `json:"page" example:"1" `                            // pagination page
	PerPage int         `json:"per_page" example:"10" `                       // pagination perpage
	Body    interface{} `json:"body" example:"{}" swaggerignore:"true"`       // Pagination bodys
	Offset  int         `json:"offset" example:"0"`
	Limit   int         `json:"limit"  example:"10" ` // Response limit
}

type VueTableResponse struct {
	Total       int         `json:"total" example:"100" `                   // Response total
	PerPage     int         `json:"per_page" example:"10" `                 // Response perpage
	CurrentPage int         `json:"current_page" example:"1" `              // Response currentpage
	LastPage    int         `json:"last_page"  example:"20" `               // Response lastpage
	NextPageURL string      `json:"next_page_url"  example:"" `             // Response nextpageurl
	PrevPageURL string      `json:"prev_page_url"  example:"" `             // Response prevpageurl
	Form        int         `json:"form" example:"0" `                      // Response form
	To          int         `json:"to"  example:"10" `                      // Response to
	Datas       interface{} `json:"datas" example:"" swaggerignore:"true" ` // Response data

}
type ResponseData struct {
	Status     int         `json:"status"`                                 // Response status
	Meta       interface{} `json:"meta" swaggerignore:"true"`              // Response meta
	Data       interface{} `json:"data" swaggerignore:"true"`              // Response data
	Success    bool        `json:"success" default:"true"  example:"true"` //Response success
	Alert      Alert       `json:"alert" swaggerignore:"true"`             // Response Alert Message
	Pagination Pagination  `json:"pagination" swaggerignore:"true"`        //Response  Pagination
}

func GetResponse() ResponseData {
	var res ResponseData
	alert := Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	res.Status = 200
	res.Meta = http.StatusText(200)
	res.Success = true
	res.Alert = alert
	res.Data = []string{}
	return res
}

func GenPagination(query url.Values, table string) Pagination {
	var p = Pagination{}
	p.Page = 1
	p.Limit = 10
	for key, value := range query {
		queryValue := value[len(value)-1]
		log.Println("key-->", key)
		switch key {
		case "per_page":
			p.Limit = limit(queryValue)
			break
		case "page":
			p.Page = page(queryValue)
			break
		case "sort":
			p.Sort = sortOrder(table, queryValue)
			break
		case "kw":
			p.Kw = queryValue
			break
		case "filter":
			p.Filter = queryValue
			break
		}
	}
	p.PerPage = p.Limit
	p.Offset = (p.Page - 1) * p.Limit
	return p
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	w.JSON(status, payload)
}

// SortOrder returns the string for sorting and orderin data
// sort="aaa|desc,bbb|asc,ccc|desc"
func sortOrder(table, sort string) string {
	arrsorts := strings.Split(sort, ",") //  --> ["aaa|desc","bbb|asc","ccc|desc"]
	var rs string
	for _, sortcol := range arrsorts {
		arsortcol := strings.Split(sortcol, "|")
		// log.Println("sort-->", arsortcol[0], arsortcol[1])
		if rs != "" {
			rs = rs + ","
		}
		rs = rs + " " + table + "." + arsortcol[0] + " " + arsortcol[1]
	}
	return rs
}

// Offset returns the starting number of result for pagination
func offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0 //default offset 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
func limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 //default limit
	}
	return limitInt
}

func page(page string) int {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1 //default limit
	}
	return pageInt
}

// Search adds where to search keywords
// filter=aaa|aaa,bbb|bbbb,ccc|ccc
func Filter(query string) map[string]string { //["aaaa"=>"bbbb","bbbb"=>"ccccc"]
	arrs := make(map[string]string)
	if query == "" {
		return arrs
	}
	arqry := strings.Split(query, ",")
	for _, v := range arqry {
		q := strings.Split(v, "|")
		log.Println("k->", q[0], "kv->", q[1])
		arrs[q[0]] = q[1]
	}
	return arrs
}

// Search adds where to search keywords
// kw=aaa bbb ccc
func Searchkw(kw string) []string {
	arrs := []string{}
	if kw == "" {
		return []string{}
	}

	kws := strings.Split(kw, " ")
	for _, v := range kws {
		arrs = append(arrs, v)
	}
	return arrs
}
