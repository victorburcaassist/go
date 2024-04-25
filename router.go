package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/items", GetItems)
	router.POST("/item", AddItem)
	router.PUT("/item", UpdateItem)
	router.DELETE("/item", DeleteItem)
}
