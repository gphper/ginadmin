/*
 * @Description:用户组管理
 * @Author: ghphper
 * @Date: 2021-07-04 11:58:45
 */

package setting

import (
	"net/http"
	"strings"

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

func NewAdminGroupController() adminGroupController {
	return adminGroupController{}
}

func (con adminGroupController) Routes(rg *gin.RouterGroup) {
	rg.GET("/index", con.index)
	rg.GET("/add", con.addIndex)
	rg.POST("/save", con.save)
	rg.GET("/edit", con.edit)
	rg.GET("/del", con.del)
}

/**
角色列表
*/
func (con adminGroupController) index(c *gin.Context) {

	var groups []string

	key := c.Query("keyword")
	if key != "" {
		for _, v := range casbinauth.GetGroups() {
			if strings.Contains(v, key) {
				groups = append(groups, key)
			}
		}
	} else {
		groups = casbinauth.GetGroups()
	}

	c.HTML(http.StatusOK, "setting/group.html", gin.H{
		"adminGroups": groups,
		"keyword":     key,
	})
}

/**
添加角色
*/
func (con adminGroupController) addIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
		"menuList": menu.GetMenu(),
		"id":       "",
	})
}

/**
保存角色
*/
func (con adminGroupController) save(c *gin.Context) {

	var req models.AdminGroupSaveReq
	err := con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	err = services.NewAdminGroupService().SaveGroup(req)
	if err != nil {
		con.Error(c, "操作失败")
		return
	}

	con.Success(c, "/admin/setting/admingroup/index", "操作成功")
}

/**
编辑
*/
func (con adminGroupController) edit(c *gin.Context) {
	id := c.Query("id")
	c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
		"menuList": menu.GetMenu(),
		"id":       id,
	})
}

/**
删除
*/
func (con adminGroupController) del(c *gin.Context) {

	id := c.Query("id")
	dbOk, dbErr := services.NewAdminGroupService().DelGroup(id)
	if dbErr != nil || !dbOk {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "", "删除成功")
	}
}
