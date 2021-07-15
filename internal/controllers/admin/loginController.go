package admin

import (
	"encoding/json"
	"ginadmin/internal/models"
	"ginadmin/pkg/comment"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginController struct {
	BaseController
}

var Lc = loginController{}

func (con *loginController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "home/login.html", gin.H{})
		} else {
			username := c.PostForm("username")
			password := c.PostForm("password")
			var adminUser models.AdminUsers
			err := models.Db.Table("admin_users").Where("username = ?", username).First(&adminUser).Error
			if err != nil {
				con.Error(c, "账号密码错误")
				return
			}
			//判断密码是否正确
			if comment.Encryption(password, adminUser.Salt) == adminUser.Password {
				//获取用户组信息
				var adminGroup models.AdminGroup
				adminGroup.GroupId = adminUser.GroupId
				err = models.Db.Find(&adminGroup).Error
				if err != nil {
					con.Error(c, "账号密码错误")
				}
				//var jsonPrivs map[string]int
				//json.Unmarshal([]byte(adminGroup.Privs), &jsonPrivs)
				userInfo := make(map[string]interface{})
				userInfo["uid"] = adminUser.Uid
				userInfo["username"] = adminUser.Username
				userInfo["groupname"] = adminGroup.GroupName
				//userInfo["privs"] = jsonPrivs
				userInfo["privs"] = adminGroup.Privs
				//session 存储数据
				session := sessions.Default(c)
				userstr, _ := json.Marshal(userInfo)

				session.Set("userInfo", string(userstr))
				session.Save()

				uio := make(map[string]interface{})
				json.Unmarshal([]byte(session.Get("userInfo").(string)), &uio)
				con.Success(c, "/admin/home", "登录成功")
			} else {
				con.Error(c, "账号密码错误")
			}

		}
	}
}

func (con *loginController) LoginOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("userInfo")
		session.Save()
		c.Redirect(http.StatusFound, "/admin/login")
	}
}
