package admin

import (
	"encoding/json"
	"ginadmin/internal/menu"
	"ginadmin/internal/models"
	"ginadmin/pkg/casbinauth"
	"ginadmin/pkg/comment"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type homeController struct {
	BaseController
}

var Hc = homeController{}

func (con *homeController) Home() gin.HandlerFunc {
	menuList := menu.GetMenu()
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userInfoJson := session.Get("userInfo")
		userData := make(map[string]interface{})
		err := json.Unmarshal([]byte(userInfoJson.(string)), &userData)
		if err != nil {
			// 取不到就是没有登录
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
			return
		}

		privs, err := casbinauth.GetGroupByUser(userData["username"].(string))
		var groupname string
		if err == nil {
			groupname = strings.Join(privs, ",")
		}

		c.HTML(http.StatusOK, "home/home.html", gin.H{
			"menuList":  menuList,
			"userInfo":  userData,
			"groupName": groupname,
		})
	}
}

func (con *homeController) Welcome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/welcome.html", gin.H{})
	}
}

func (con *homeController) EditPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		c.HTML(http.StatusOK, "home/password_form.html", gin.H{
			"id": id,
		})
	}
}

func (con *homeController) SavePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		oldPassword := c.PostForm("old_password")
		newPassword := c.PostForm("new_password")
		subPassword := c.PostForm("sub_password")

		if newPassword != subPassword {
			con.Error(c, "请确认新密码")
			return
		}

		adminUser, _ := models.GetAdminUserById(id)
		oldPass := comment.Encryption(oldPassword, adminUser.Salt)
		if oldPass != adminUser.Password {
			con.Error(c, "原密码不正确")
			return
		}

		newPass := comment.Encryption(newPassword, adminUser.Salt)
		err := models.AlterAdminUserPass(id, newPass)
		if err != nil {
			con.Error(c, "修改密码失败")
			return
		} else {
			con.Success(c, "", "修改成功")
			return
		}
	}
}
