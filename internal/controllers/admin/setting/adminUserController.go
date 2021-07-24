package setting

import (
	"ginadmin/internal/controllers/admin"
	"ginadmin/internal/models"
	"ginadmin/internal/services"
	"ginadmin/pkg/casbinauth"
	"ginadmin/pkg/comment"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminUserController struct {
	admin.BaseController
}

var Auc = adminUserController{}

/**
管理员列表
*/
func (con *adminUserController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err           error
			req           models.AdminUserIndexReq
			adminUserList []models.AdminUsers
		)

		err = con.FormBind(c, &req)
		if err != nil {
			con.Error(c, err.Error())
			return
		}

		adminDb := services.AuService.GetAdminUsers(req)
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
func (con *adminUserController) AddIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "setting/adminuser_form.html", gin.H{
			"adminGroups": casbinauth.GetGroups(),
		})
	}
}

/**
保存
*/
func (con *adminUserController) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req models.AdminUserSaveReq
		)

		err = con.FormBind(c, &req)
		if err != nil {
			con.Error(c, err.Error())
			return
		}
		err = services.AuService.SaveAdminUser(req)
		if err == nil {
			con.Success(c, "/admin/setting/adminuser/index", "操作成功")
		} else {
			con.Error(c, err.Error())
		}
	}
}

/**
编辑
*/
func (con *adminUserController) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		adminGroup, _ := services.AgService.GetAllGroup()
		adminUser, _ := services.AuService.GetAdminUser(id)

		c.HTML(http.StatusOK, "setting/adminuser_form.html", gin.H{
			"adminGroups": adminGroup,
			"adminUser":   adminUser,
		})
	}
}

/**
删除
*/
func (con *adminUserController) Del() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		err := services.AuService.DelAdminUser(id)
		if err != nil {
			con.Error(c, "删除失败")
		} else {
			con.Success(c, "", "删除成功")
		}

	}
}
