/*
 * @Description:后台主页
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gphper/ginadmin/internal/menu"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/casbinauth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type homeController struct {
	BaseController
}

var Hc = homeController{}

func (con homeController) Home(c *gin.Context) {
	menuList := menu.GetMenu()

	session := sessions.Default(c)
	userInfoJson := session.Get("userInfo")
	userData := make(map[string]interface{})
	err := json.Unmarshal([]byte(userInfoJson.(string)), &userData)
	if err != nil {
		// 取不到就是没有登录
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">top.location.href="/admin/login"</script>`)
		return
	}

	privs, err := casbinauth.GetGroupByUser(userData["username"].(string))
	var groupname string
	if err == nil {
		groupname = strings.Join(privs, ",")
	}

	//获取当前用户的皮肤
	uid, _ := userData["uid"].(float64)

	adminUser, _ := services.AuService.GetAdminUser(strconv.Itoa(int(uid)))

	c.HTML(http.StatusOK, "home/home.html", gin.H{
		"menuList":  menuList,
		"userInfo":  userData,
		"groupName": groupname,
		"header":    adminUser.Header,
		"logo":      adminUser.Logo,
		"sider":     adminUser.Side,
	})
}

func (con homeController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "home/welcome.html", gin.H{})
}

func (con homeController) EditPassword(c *gin.Context) {
	id := c.Query("id")
	c.HTML(http.StatusOK, "home/password_form.html", gin.H{
		"id": id,
	})
}

func (con homeController) SavePassword(c *gin.Context) {
	var (
		req models.AdminUserEditPassReq
		err error
	)
	con.FormBind(c, &req)
	err = services.AuService.EditPass(req)

	if err != nil {
		con.Error(c, err.Error())
		return
	} else {
		con.Success(c, "", "修改成功")
		return
	}
}

func (con homeController) SaveSkin(c *gin.Context) {

	var skinReq models.AdminUserSkinReq

	con.FormBind(c, &skinReq)

	value, _ := c.Get("userInfo")

	userData := make(map[string]interface{})

	fmt.Println(value.(string))

	json.Unmarshal([]byte(value.(string)), &userData)

	v, _ := userData["uid"].(float64)

	skinReq.Uid = int(v)

	services.AuService.EditSkin(skinReq)

	c.JSON(http.StatusOK, skinReq)
}
