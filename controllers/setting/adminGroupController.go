package setting

import (
	"encoding/json"
	"ginadmin/comment/menu"
	"ginadmin/controllers"
	"ginadmin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminGroupController struct {
	controllers.BaseController
}

/**
角色列表
*/
func (con *AdminGroupController) Index() gin.HandlerFunc {
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
func (con *AdminGroupController) AddIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "setting/group_form.html", gin.H{
			"menuList": menu.GetMenu(),
		})
	}
}

/**
保存角色
*/
func (con *AdminGroupController) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var privsJsonStr string
		privMap := make(map[string]struct{})
		privs, _ := c.GetPostFormArray("privs[]")
		//将数组转为map便于提高后面的判断效率
		for _, v := range privs {
			privMap[v] = struct{}{}
		}

		privsJson, err := json.Marshal(privMap)
		if err == nil {
			privsJsonStr = string(privsJson)
		} else {
			privsJsonStr = `[]`
		}

		groupName := c.PostForm("groupname")
		groupId, err := strconv.Atoi(c.PostForm("groupid"))
		if err != nil {
			groupId = 0
		}

		dbErr := models.SaveAdminGroup(uint(groupId), groupName, privsJsonStr)
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
func (con *AdminGroupController) Edit() gin.HandlerFunc {
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
func (con *AdminGroupController) Del() gin.HandlerFunc {
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
