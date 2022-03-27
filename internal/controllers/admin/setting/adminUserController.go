/*
 * @Description:用户管理
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package setting

import (
	"encoding/json"
	"net/http"

	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/casbinauth"
	"github.com/gphper/ginadmin/pkg/paginater"

	"github.com/gin-gonic/gin"
)

type adminUserController struct {
	admin.BaseController
}

var Auc = adminUserController{}

/**
管理员列表
*/
func (con adminUserController) Index(c *gin.Context) {
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
	adminUserData := paginater.PageOperation(c, adminDb, 1, &adminUserList)
	c.HTML(http.StatusOK, "setting/adminuser.html", gin.H{
		"adminUserData": adminUserData,
		"created_at":    c.Query("created_at"),
		"nickname":      c.Query("nickname"),
	})

}

/**
添加
*/
func (con adminUserController) AddIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "setting/adminuser_form.html", gin.H{
		"adminGroups": casbinauth.GetGroups(),
	})
}

/**
保存
*/
func (con adminUserController) Save(c *gin.Context) {

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

/**
编辑
*/
func (con adminUserController) Edit(c *gin.Context) {
	id := c.Query("id")
	adminUser, _ := services.AuService.GetAdminUser(id)
	var groupName []string
	json.Unmarshal([]byte(adminUser.GroupName), &groupName)
	var groupMap = make(map[string]struct{})
	for _, v := range groupName {
		groupMap[v] = struct{}{}
	}
	c.HTML(http.StatusOK, "setting/adminuser_form.html", gin.H{
		"adminGroups": casbinauth.GetGroups(),
		"adminUser":   adminUser,
		"groupMap":    groupMap,
	})
}

/**
删除
*/
func (con adminUserController) Del(c *gin.Context) {
	id := c.Query("id")
	err := services.AuService.DelAdminUser(id)
	if err != nil {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "", "删除成功")
	}

}
