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
// @response 200 {object} models.Product "OK"
// @router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
