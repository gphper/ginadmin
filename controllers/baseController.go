package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (Base *BaseController) Success(c *gin.Context, url string, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"msg":         message,
		"url":         url,
		"iframe_jump": false,
	})
}

func (Base *BaseController) Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": false,
		"msg":    message,
	})
}
