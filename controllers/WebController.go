package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Message godoc
// @summary Message
// @description Message for the service
// @id Message
// plain json xml x-www-form-urlencoded
// @produce plain
// @response 200 {object} models.Transaction "OK"
// get post put delete patch
// @router /message [get]
func Message(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// Ping godoc
// @summary Ping
// @description Ping for the service
// @id Ping
// @produce plain
// @response 200 {object} models.Product "OK"
// @router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Test godoc
// @summary Test
// @description Test for the service
// @id Test
// @produce plain
// @response 200 {object} models.User "OK"
// @router /test [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
