package setting

import (
	"encoding/json"
	"ginadmin/internal/controllers/admin"
	"ginadmin/internal/menu"
	"ginadmin/internal/models"
	"net/http"
	"strconv"

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
		adminGroups, _ := models.GetAllAdminGroup()
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

		type groupForm struct {
			Privs     []string `form:"privs[]" label:"权限" json:"privs" binding:"required"`
			GroupName string   `form:"groupname" label:"用户组名" json:"groupname" binding:"required"`
		}
		var gf groupForm

		err := con.FormBind(c, &gf)
		if err != nil {
			con.Error(c, err.Error())
			return
		}

		var privsJsonStr string
		privMap := make(map[string]struct{})
		//将数组转为map便于提高后面的判断效率
		for _, v := range gf.Privs {
			privMap[v] = struct{}{}
		}

		privsJson, err := json.Marshal(privMap)
		if err == nil {
			privsJsonStr = string(privsJson)
		} else {
			privsJsonStr = `[]`
		}

		groupId, err := strconv.Atoi(c.PostForm("groupid"))
		if err != nil {
			groupId = 0
		}

		dbErr := models.SaveAdminGroup(uint(groupId), gf.GroupName, privsJsonStr)
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
		adminGroup, _ := models.FindAdminGroupById(id)
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
		models.Db.Where("group_id = ?", id).Delete(models.AdminGroup{})
		dbErr := models.DelAdminGroupById(id)
		if dbErr != nil {
			con.Error(c, "删除失败")
		} else {
			con.Success(c, "", "删除成功")
		}
	}
}
