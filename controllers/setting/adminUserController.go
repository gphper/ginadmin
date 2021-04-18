package setting

import (
	"ginadmin/comment"
	"ginadmin/controllers"
	"ginadmin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminUserController struct {
	controllers.BaseController
}

/**
管理员列表
*/
func (this *AdminUserController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Query("p")
		page, _ := strconv.Atoi(p)

		type adminUser struct {
			models.AdminUsers
			GroupName string
		}
		var adminUserList []adminUser
		adminDb := models.Db.Table("admin_users").Joins("join admin_groups on admin_groups.group_id = admin_users.group_id").Select("admin_users.*,admin_groups.group_name").Where("uid != ?", 1)
		adminUserData := comment.PageOperation(c, adminDb, 1, page, &adminUserList)
		c.HTML(http.StatusOK, "adminuser.html", gin.H{
			"adminUserData": adminUserData,
			"created_at":    c.Query("created_at"),
			"nickname":      c.Query("nickname"),
		})
	}
}

/**
添加
*/
func (this *AdminUserController) AddIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取角色
		var adminGroup []models.AdminGroup
		models.Db.Where("group_id != ?", 1).Find(&adminGroup)
		c.HTML(http.StatusOK, "adminuser_form.html", gin.H{
			"adminGroups": adminGroup,
		})
	}
}

/**
保存
*/
func (this *AdminUserController) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		username, ubool := c.GetPostForm("username")
		password, pbool := c.GetPostForm("password")
		nickname := c.PostForm("nickname")
		phone := c.PostForm("phone")
		groupid := c.PostForm("groupid")
		groupidd, _ := strconv.Atoi(groupid)
		uid := c.PostForm("uid")
		uidd, _ := strconv.Atoi(uid)
		adminUser := models.AdminUsers{
			Uid:       uint(uidd),
			GroupId:   uint(groupidd),
			Username:  "",
			Nickname:  nickname,
			Password:  "",
			Phone:     phone,
			LastLogin: "",
			Salt:      "",
			ApiToken:  "",
		}

		if ubool {
			adminUser.Username = username
		}

		if pbool && password != "" {
			salt := comment.RandString(6)
			adminUser.Salt = salt
			passwordSalt := comment.Encryption(password, salt)
			adminUser.Password = passwordSalt
		}
		if uidd > 0 {
			err = models.Db.Model(&adminUser).Update(adminUser).Error
		} else {
			err = models.Db.Save(&adminUser).Error
		}

		if err == nil {
			this.Success(c, "/admin/setting/adminuser/index", "操作成功")
		} else {
			this.Error(c, "操作失败")
		}
	}
}

/**
编辑
*/
func (this *AdminUserController) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		var adminGroup []models.AdminGroup
		models.Db.Find(&adminGroup)
		var adminUser models.AdminUsers
		models.Db.Where("uid = ?", id).First(&adminUser)
		c.HTML(http.StatusOK, "adminuser_form.html", gin.H{
			"adminGroups": adminGroup,
			"adminUser":   adminUser,
		})
	}
}

/**
删除
*/
func (this *AdminUserController) Del() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		models.Db.Where("uid = ?", id).Delete(models.AdminUsers{})
		this.Success(c, "", "删除成功")
	}
}
