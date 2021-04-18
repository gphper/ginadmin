package middleware

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminUserAuth() gin.HandlerFunc{
	return func(c *gin.Context){
		session := sessions.Default(c)
		userInfoJson := session.Get("userInfo")
		if userInfoJson == nil{
			// 取不到就是没有登录
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
			return
		}
		userData := make(map[string]interface{})
		err := json.Unmarshal([]byte(userInfoJson.(string)),&userData)
		if err != nil{
			// 取不到就是没有登录
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
			return
		}else{
			uri := c.FullPath()
			userPrivs := userData["privs"]
			userPrivsSlice := userPrivs.(map[string]interface{})
			//将url转为index
			_,ook := userPrivsSlice["all"]
			_,okk := userPrivsSlice[uri]
			if ook || okk || uri == "/admin/home/"{
				c.Next()
			}else{
				c.JSON(http.StatusNonAuthoritativeInfo,gin.H{
					"Status":"无权限禁止访问",
				})
				c.Abort()
			}
		}
	}
}
