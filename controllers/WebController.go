package controllers

import (
	"apitest/db"
	"apitest/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Message godoc
// @summary Message
// @description Message for the service
// @id Message
// plain json xml x-www-form-urlencoded
// @produce plain
// @response 200 {object} utils.ResponseData "OK"
// get post put delete patch
// @router / [get]
func Message(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to WebApi V0.0.1",
	})
}

// Ping godoc
// @summary Ping
// @description Ping for the service
// @id Ping
// @produce plain
// @response 200 {object} utils.ResponseData  "OK"
// @router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Ping godoc
// @summary Ping
// @description  test User Relagion with user->role->permission
// @id Test
// @produce plain
// @response 200 {object} utils.ResponseData  "OK"
// @router /test [get]
func Test(c *gin.Context) {
	var user models.User
	if err := db.GetDB().Where("id = ? ", 1).Preload("Roles.Permissions").Find(&user).Error; err != nil {
		log.Fatal(err)
		return
	}
	log.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
