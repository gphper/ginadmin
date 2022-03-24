/*
 * @Description:用户组管理
 * @Author: ghphper
 * @Date: 2021-07-04 11:58:45
 */

package setting

import (
	"net/http"

	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/menu"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/casbinauth"

	"github.com/gin-gonic/gin"
)

type adminGroupController struct {
	admin.BaseController
}

var Agc = adminGroupController{}

/**
角色列表
*/
func (con adminGroupController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "setting/group.html", gin.H{
		"adminGroups": casbinauth.GetGroups(),
	})
}

/**
添加角色
*/
func (con adminGroupController) AddIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
		"menuList": menu.GetMenu(),
		"id":       "",
	})
}

/**
保存角色
*/
func (con adminGroupController) Save(c *gin.Context) {

	var req models.AdminGroupSaveReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	dbErr := services.AgService.SaveGroup(req)
	if dbErr != nil {
		con.Error(c, "操作失败")
	} else {
		con.Success(c, "/admin/setting/admingroup/index", "操作成功")
	}
}

/**
编辑
*/
func (con adminGroupController) Edit(c *gin.Context) {
	id := c.Query("id")
	c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
		"menuList": menu.GetMenu(),
		"id":       id,
	})
}

/**
删除
*/
func (con adminGroupController) Del(c *gin.Context) {

	id := c.Query("id")
	dbOk, dbErr := services.AgService.DelGroup(id)
	if dbErr != nil || !dbOk {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "", "删除成功")
	}
}
