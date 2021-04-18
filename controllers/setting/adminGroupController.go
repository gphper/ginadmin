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
func(con *AdminGroupController) Index() gin.HandlerFunc{
	return func(c *gin.Context) {
		var adminGroups []models.AdminGroup
		models.Db.Where("group_id != ?",1).Find(&adminGroups)
		c.HTML(http.StatusOK,"group.html",gin.H{
			"adminGroups":adminGroups,
		})
	}
}

/**
添加角色
 */
func(con *AdminGroupController) AddIndex() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.HTML(http.StatusOK,"group_form.html",gin.H{
			"menuList":menu.GetMenu(),
		})
	}
}

/**
保存角色
 */
func(con *AdminGroupController) Save() gin.HandlerFunc{
	return func(c *gin.Context) {
		var privsJsonStr string
		privMap := make(map[string]int)
		privs,_ := c.GetPostFormArray("privs[]")
		//将数组转为map便于提高后面的判断效率
		for k,v := range privs{
			privMap[v] = k
		}

		privsJson,err := json.Marshal(privMap)
		if err == nil {
			privsJsonStr = string(privsJson)
		}else{
			privsJsonStr = `[]`
		}

		groupname := c.PostForm("groupname")
		groupid,err := strconv.Atoi(c.PostForm("groupid"))
		if err != nil {
			groupid = 0
		}

		adminGroup := models.AdminGroup{
			GroupId:   uint(groupid),
			GroupName: groupname,
			Privs:     privsJsonStr,
		}

		models.Db.Save(&adminGroup)
		con.Success(c,"/admin/setting/admingroup/index","操作成功")
	}
}

/**
编辑
 */
func(con *AdminGroupController) Edit() gin.HandlerFunc{
	return func(c *gin.Context) {
		var adminGroup models.AdminGroup
		id := c.Query("id")
		models.Db.Where("group_id = ?",id).First(&adminGroup)
		var jsonPrivs map[string]interface{}
		json.Unmarshal([]byte(adminGroup.Privs),&jsonPrivs)
		c.HTML(http.StatusOK,"group_form.html",gin.H{
			"adminGroup":adminGroup,
			"jsonPrivs": jsonPrivs,
			"menuList":menu.GetMenu(),
		})
	}
}

/**
删除
 */
func(con *AdminGroupController) Del() gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		models.Db.Where("group_id = ?",id).Delete(models.AdminGroup{})
		con.Success(c,"","删除成功")
	}
}