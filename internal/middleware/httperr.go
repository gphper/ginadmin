/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-09 19:48:41
 */
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotHttpStatusOk() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if c.Writer.Status() > 399 {
			if ok := strings.Contains(c.Request.RequestURI, "/api"); ok {
				//TODO 接口处理

			} else {
				//页面处理
				c.HTML(http.StatusOK, "home/error.html", gin.H{
					"title": c.Writer.Status(),
				})
			}
		}
	}
}
