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

	if c.GetHeader("Accept") == "application/json" {

		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "url not fund",
			"data": "",
		})
	} else {

		c.HTML(http.StatusOK, "home/error.html", gin.H{
			"title": "出错了~",
			"code":  404,
			"msg":   "url not fund",
		})
	}
}
