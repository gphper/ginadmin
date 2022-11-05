/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-18 19:51:55
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handleController struct {
}

func NewHandleController() handleController {
	return handleController{}
}

func (con handleController) Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "home/error.html", gin.H{
		"title": c.Writer.Status(),
	})
}
