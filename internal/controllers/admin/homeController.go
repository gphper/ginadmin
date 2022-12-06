/*
 * @Description:后台主页
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package admin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gphper/ginadmin/internal/menu"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/casbinauth"

	"github.com/gin-gonic/gin"
)

type homeController struct {
	BaseController
}

func NewHomeController() homeController {
	return homeController{}
}

func (con homeController) Routes(rg *gin.RouterGroup) {
	rg.GET("/", con.home)
	rg.GET("/welcome", con.welcome)
	rg.GET("/edit_password", con.editPassword)
	rg.POST("/save_password", con.savePassword)
	rg.POST("/save_skin", con.saveSkin)
}

func (con homeController) home(c *gin.Context) {
	menuList := menu.GetMenu()

	username, ok := c.Get("username")
	if !ok {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
		return
	}
	privs, err := casbinauth.GetGroupByUser(username.(string))
	var groupname string
	if err == nil {
		groupname = strings.Join(privs, ",")
	}

	//获取当前用户的皮肤
	uid, ok := c.Get("uid")
	if !ok {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
		return
	}
	adminUser, _ := (services.NewAdminUserService()).GetAdminUser(map[string]interface{}{"uid": uid})

	con.Html(c, http.StatusOK, "home/home.html", gin.H{
		"menuList":  menuList,
		"groupName": groupname,
		"header":    adminUser.Header,
		"logo":      adminUser.Logo,
		"sider":     adminUser.Side,
	})
}

func (con homeController) welcome(c *gin.Context) {
	con.Html(c, http.StatusOK, "home/welcome.html", gin.H{})
}

func (con homeController) editPassword(c *gin.Context) {
	id := c.Query("id")
	con.Html(c, http.StatusOK, "home/password_form.html", gin.H{
		"id": id,
	})
}

/**
*	修改密码
 */
func (con homeController) savePassword(c *gin.Context) {

	var (
		req models.AdminUserEditPassReq
		err error
	)
	con.FormBind(c, &req)

	err = services.NewAdminUserService().EditPass(req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, "", "修改成功")
}

/**
*	保存皮肤
 */
func (con homeController) saveSkin(c *gin.Context) {

	var skinReq models.AdminUserSkinReq

	con.FormBind(c, &skinReq)

	uid, _ := c.Get("uid")
	fmt.Printf("%T", uid)
	skinReq.Uid = uid.(int)

	err := services.NewAdminUserService().EditSkin(skinReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, skinReq)
}
