package utils

import (
	"log"
	"net/http"

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

type ResponseData struct {
	Status  int         `json:status`                                    // Response status
	Meta    interface{} `json:"meta" swaggerignore:"true"`               // Response meta
	Data    interface{} `json:"data" swaggerignore:"true"`               // Response data
	Success bool        `json:"success" default: "true"  example:"true"` //Response success
	Msg     string      `json:"msg"`
	Title   string      `json:"Title" example:"Success" ` // Response type
	Type    string      `json:"type" example:"success" `  // Response type

}

func RespondJSON(w *gin.Context, status int, payload interface{}, alert Alert) {
	log.Println("status ", status)
	var res ResponseData
	res.Meta = http.StatusText(status)
	res.Status = status
	if status == 200 {
		res.Success = true
		res.Title = "Success"
		res.Type = "success"
	} else {
		res.Title = "Error"
		res.Type = "Error"
	}
	res.Msg = http.StatusText(status)
	res.Data = payload
	w.JSON(res.Status, gin.H{
		"status":  res.Status,
		"meta":    res.Meta,
		"data":    res.Data,
		"success": res.Success,
		"type":    alert.Type,
		"title":   alert.Title,
		"msg":     alert.Msg,
		"alert":   alert,
	})
}
