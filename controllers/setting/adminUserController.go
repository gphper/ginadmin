package setting

import (
	"ginadmin/comment"
	"ginadmin/controllers"
	"ginadmin/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	controllers.BaseController
}

/**
管理员列表
*/
func (this *AdminUserController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		nickname := c.Query("nickname")
		createdAt := c.Query("created_at")

		type adminUser struct {
			models.AdminUsers
			GroupName string
		}
		var adminUserList []adminUser

		adminDb := models.GetAllAdminUserJoinGroup()

		if nickname != "" {
			adminDb = models.GetAllAdminUserJoinGroupLikeNickname(adminDb, nickname)
		}

		if createdAt != "" {
			period := strings.Split(createdAt, " ~ ")
			start := period[0] + " 00:00:00"
			end := period[1] + " 23:59:59"
			adminDb = models.GetAllAdminUserJoinGroupTimeRange(adminDb, start, end)
		}

		adminUserData := comment.PageOperation(c, adminDb, 1, &adminUserList)
		c.HTML(http.StatusOK, "setting/adminuser.html", gin.H{
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
		adminGroups, _ := models.GetAllAdminGroup()
		c.HTML(http.StatusOK, "setting/adminuser_form.html", gin.H{
			"adminGroups": adminGroups,
		})
	}
}

/**
保存
*/
func (this *AdminUserController) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		username, _ := c.GetPostForm("username")
		password, _ := c.GetPostForm("password")
		nickname := c.PostForm("nickname")
		phone := c.PostForm("phone")
		groupid := c.PostForm("groupid")
		groupidd, _ := strconv.Atoi(groupid)
		uid := c.PostForm("uid")
		uidd, _ := strconv.Atoi(uid)
		if uidd > 0 {
			err = models.SaveAdminUser(uidd, groupidd, nickname, phone, password)
		} else {
			err = models.AddAdminUser(groupidd, username, nickname, phone, password)
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
		c.HTML(http.StatusOK, "setting/adminuser_form.html", gin.H{
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
		id := c.Query("id")
		models.Db.Where("uid = ?", id).Delete(models.AdminUsers{})
		this.Success(c, "", "删除成功")
	}
}
