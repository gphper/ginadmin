/*
 * @Description:用户管理
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package setting

import (
	"context"
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

func NewAdminUserController() adminUserController {
	return adminUserController{}
}

func (con adminUserController) Routes(rg *gin.RouterGroup) {
	rg.GET("/index", con.index)
	rg.GET("/add", con.addIndex)
	rg.POST("/save", con.save)
	rg.GET("/edit", con.edit)
	rg.GET("/del", con.del)
}

/**
管理员列表
*/
func (con adminUserController) index(c *gin.Context) {
	var (
		err           error
		req           models.AdminUserIndexReq
		adminUserList []models.AdminUsers
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.ErrorHtml(c, err)
		return
	}

	ctx, _ := c.Get("ctx")

	adminDb := services.NewAdminUserService().GetAdminUsers(ctx.(context.Context), req)
	adminUserData, err := paginater.PageOperation(c, adminDb, 1, &adminUserList)
	if err != nil {
		con.ErrorHtml(c, err)
	}
	con.Html(c, http.StatusOK, "setting/adminuser.html", gin.H{
		"adminUserData": adminUserData,
		"created_at":    c.Query("created_at"),
		"nickname":      c.Query("nickname"),
	})

}

/**
添加
*/
func (con adminUserController) addIndex(c *gin.Context) {
	con.Html(c, http.StatusOK, "setting/adminuser_form.html", gin.H{
		"adminGroups": casbinauth.GetGroups(),
	})
}

/**
保存
*/
func (con adminUserController) save(c *gin.Context) {

	var (
		err error
		req models.AdminUserSaveReq
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	err = services.NewAdminUserService().SaveAdminUser(req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, "/admin/setting/adminuser/index", "操作成功")
}

/**
编辑
*/
func (con adminUserController) edit(c *gin.Context) {
	id := c.Query("id")
	adminUser, _ := services.NewAdminUserService().GetAdminUser(map[string]interface{}{"uid": id})
	var groupName []string
	json.Unmarshal([]byte(adminUser.GroupName), &groupName)
	var groupMap = make(map[string]struct{})
	for _, v := range groupName {
		groupMap[v] = struct{}{}
	}
	con.Html(c, http.StatusOK, "setting/adminuser_form.html", gin.H{
		"adminGroups": casbinauth.GetGroups(),
		"adminUser":   adminUser,
		"groupMap":    groupMap,
	})
}

/**
删除
*/
func (con adminUserController) del(c *gin.Context) {

	id := c.Query("id")

	err := services.NewAdminUserService().DelAdminUser(id)
	if err != nil {
		con.Error(c, "删除失败")
		return
	}

	con.Success(c, "", "删除成功")
}
