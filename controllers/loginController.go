package controllers

import (
	"encoding/json"
	"ginadmin/comment"
	"ginadmin/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "home/login.html", gin.H{})
		} else {
			username := c.PostForm("username")
			password := c.PostForm("password")
			var adminUser models.AdminUsers
			err := models.Db.Table("admin_users").Where("username = ?", username).First(&adminUser).Error
			if err != nil {
				this.Error(c, "用户不存在")
				return
			}

			//判断密码是否正确
			if comment.Encryption(password, adminUser.Salt) == adminUser.Password {
				//获取用户组信息
				var adminGroup models.AdminGroup
				adminGroup.GroupId = adminUser.GroupId
				err = models.Db.Find(&adminGroup).Error
				if err != nil {
					this.Error(c, "用户不存在")
				}
				var jsonPrivs map[string]int
				json.Unmarshal([]byte(adminGroup.Privs), &jsonPrivs)
				userInfo := make(map[string]interface{})
				userInfo["uid"] = adminUser.Uid
				userInfo["username"] = adminUser.Username
				userInfo["groupname"] = adminGroup.GroupName
				userInfo["privs"] = jsonPrivs
				//session 存储数据
				session := sessions.Default(c)
				userstr, _ := json.Marshal(userInfo)
				session.Set("userInfo", string(userstr))
				session.Save()
				c.Redirect(http.StatusFound, "/admin/home")
			} else {
				this.Error(c, "账号密码错误")
			}

		}
	}
}

func (this *LoginController) LoginOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("userInfo")
		session.Save()
		c.Redirect(http.StatusFound, "/admin/login")
	}
}
