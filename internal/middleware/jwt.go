/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-26 19:56:28
 */
package middleware

import (
	"net/http"

	"github.com/gphper/ginadmin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		jtoken := c.Request.Header.Get("Authorization")
		payload, err := jwt.Check(jtoken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
		}
		c.Set("uid", payload.Uid)
		c.Next()
	}
}
