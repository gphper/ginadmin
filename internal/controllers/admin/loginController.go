/*
 * @Description:后台登录相关方法
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package admin

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/captcha/store"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type loginController struct {
	BaseController
}

func NewLoginController() loginController {
	return loginController{}
}

func (con loginController) Routes(rg *gin.RouterGroup) {
	rg.GET("/captcha", con.captcha)
	/*******登录路由**********/
	rg.GET("/login", con.login)
	rg.POST("/login", con.login)
	rg.GET("/login_out", con.loginOut)
	rg.POST("/login_out", con.loginOut)
}

/**
* 登录
 */
func (con loginController) login(c *gin.Context) {
	if c.Request.Method == "GET" {
		con.Html(c, http.StatusOK, "home/login.html", gin.H{
			"title": "GinAdmin管理平台",
		})
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 为测试方便release模式才开启验证码
		if gin.Mode() == gin.ReleaseMode {

			captch := c.PostForm("captcha")
			var store = store.NewSessionStore(c, 20)
			verify := store.Verify("", captch, true)
			if !verify {
				con.Error(c, "验证码错误")
				return
			}

		}

		adminUser, err := services.NewAdminUserService().GetAdminUser(map[string]interface{}{"username": username})
		if err != nil {
			con.Error(c, "账号密码错误")
			return
		}
		//判断密码是否正确
		if gstrings.Encryption(password, adminUser.Salt) == adminUser.Password {

			userInfo := make(map[string]interface{})
			userInfo["uid"] = adminUser.Uid
			userInfo["username"] = adminUser.Username
			userInfo["groupname"] = adminUser.GroupName
			//session 存储数据
			session := sessions.Default(c)
			userstr, _ := json.Marshal(userInfo)

			session.Set("userInfo", string(userstr))
			session.Save()

			con.Success(c, "/admin/home", "登录成功")
		} else {
			con.Error(c, "账号密码错误")
		}

	}

}

/**
* 登出
 */
func (con loginController) loginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userInfo")
	session.Save()
	c.Redirect(http.StatusFound, "/admin/login")
}

/*
* 验证码
 */
func (con loginController) captcha(c *gin.Context) {

	var store = store.NewSessionStore(c, 20)
	driver := &base64Captcha.DriverString{
		Height: 60,
		Width:  150,
		Length: 4,
		Source: "abcdefghijklmnopqr234509867",
	}
	draw := base64Captcha.NewCaptcha(driver, store)
	_, b64s, err := draw.Generate()
	if err != nil {
		con.Error(c, "获取验证码失败")
	}

	i := strings.Index(b64s, ",")
	if i < 0 {
		log.Fatal("no comma")
	}
	// pass reader to NewDecoder
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64s[i+1:]))

	io.Copy(c.Writer, dec)
}
