package setting

import (
	"encoding/json"
	"ginadmin/internal/controllers/admin"
	"ginadmin/internal/menu"
	"ginadmin/internal/models"
	"ginadmin/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminGroupController struct {
	admin.BaseController
}

var Agc = adminGroupController{}

/**
角色列表
*/
func (con *adminGroupController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminGroups, _ := services.AgService.GetList()
		c.HTML(http.StatusOK, "setting/group.html", gin.H{
			"adminGroups": adminGroups,
		})
	}
}

/**
添加角色
*/
func (con *adminGroupController) AddIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
			"menuList": menu.GetMenu(),
		})
	}
}

/**
保存角色
*/
func (con *adminGroupController) Save() gin.HandlerFunc {
	return func(c *gin.Context) {

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
}

/**
编辑
*/
func (con *adminGroupController) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		adminGroup, _ := services.AgService.GetGroup(id)
		var jsonPrivs map[string]struct{}
		json.Unmarshal([]byte(adminGroup.Privs), &jsonPrivs)
		c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
			"adminGroup": adminGroup,
			"jsonPrivs":  jsonPrivs,
			"menuList":   menu.GetMenu(),
		})
	}
}

/**
删除
*/
func (con *adminGroupController) Del() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		dbErr := services.AgService.DelGroup(id)
		if dbErr != nil {
			con.Error(c, "删除失败")
		} else {
			con.Success(c, "", "删除成功")
		}
	}
}
