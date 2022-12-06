/*
 * @Description:后台管理中间件
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */
package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gphper/ginadmin/pkg/casbinauth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/**
用户登录验证
*/
func AdminUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		type UserData struct {
			Uid       int    `json:"uid"`
			Username  string `json:"username"`
			GroupName string `json:"groupname"`
		}
		var userData UserData

		session := sessions.Default(c)
		userInfoJson := session.Get("userInfo")
		if userInfoJson == nil {
			// 取不到就是没有登录
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
			return
		}

		err := json.Unmarshal([]byte(userInfoJson.(string)), &userData)
		if err != nil {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
			return
		}

		c.Set("username", userData.Username)
		c.Set("uid", userData.Uid)
		c.Set("groupname", userData.GroupName)
		c.Next()
	}
}

/**
用户权限验证
*/
func AdminUserPrivs() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userInfoJson := session.Get("userInfo")
		userData := make(map[string]interface{})
		err := json.Unmarshal([]byte(userInfoJson.(string)), &userData)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"msg":    "无权限禁止操作",
			})

			c.Abort()
		}
		uri := c.FullPath()
		ok, err := casbinauth.Check(userData["username"].(string), uri, c.Request.Method)
		if !ok || err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"msg":    "无权限禁止操作",
			})
			c.Abort()
		}
		c.Next()
	}
}
