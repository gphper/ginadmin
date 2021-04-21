package controllers

import (
	"encoding/json"
	"ginadmin/comment/menu"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
	BaseController
}

func (con *HomeController)Home() gin.HandlerFunc {
	menuList := menu.GetMenu()
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userInfoJson := session.Get("userInfo")
		userData := make(map[string]interface{})
		err := json.Unmarshal([]byte(userInfoJson.(string)),&userData)
		if err != nil{
			// 取不到就是没有登录
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
			return
		}
		c.HTML(http.StatusOK,"home/home.html",gin.H{
			"menuList":menuList,
			"userInfo":userData,
			"userPrivs":userData["privs"],
		})
	}
}