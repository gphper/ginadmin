/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-31 19:59:19
 */
package setting

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/pkg/mysqlx"
	"github.com/gphper/ginadmin/pkg/redisx"
	"github.com/gphper/ginadmin/pkg/utils/httptestutil"
	"github.com/gphper/ginadmin/pkg/utils/httptestutil/router"
	"github.com/gphper/ginadmin/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AdminTestSuite struct {
	suite.Suite
	router  *router.Router
	cookies []*http.Cookie
}

func (suite *AdminTestSuite) SetupSuite() {

	router, err := router.Init()
	if err != nil {
		log.Fatalf("router init fail: %s", err)
	}

	router.SetAdminRoute("/admin", admin.NewLoginController())
	router.SetAdminRoute("/admin/setting/admingroup", NewAdminGroupController())

	suite.router = router
}

// 登录页测试
func (suite *AdminTestSuite) TestALoginGet() {
	body := httptestutil.Get("/admin/login", suite.router)
	assert.Contains(suite.T(), string(body), "登录")
}

// 登录
func (suite *AdminTestSuite) TestALoginPost() {
	option := httptestutil.OptionValue{
		Param: url.Values{
			"username": {"admin"},
			"password": {"111111"},
		},
	}
	// 发起post请求，以表单形式传递参数
	body, cookies := httptestutil.PostForm("/admin/login", suite.router, option)
	if len(cookies) != 0 {
		suite.cookies = cookies
	}

	assert.Contains(suite.T(), string(body), `"status":true`)
}

// 添加角色
func (suite *AdminTestSuite) TestBAddGroup() {
	param := httptestutil.OptionValue{
		Param: url.Values{
			"groupname": {"测试角色11"},
			"privs[]": []string{
				"setting:get",
				"/admin/setting/adminuser/index:get",
				"/admin/setting/adminuser/add:get",
				"/admin/setting/adminuser/edit:get",
				"/admin/setting/adminuser/save:post",
				"/admin/setting/admingroup/index:get",
				"/admin/setting/admingroup/add:get",
				"/admin/setting/admingroup/edit:get",
				"/admin/setting/admingroup/save:post",
				"/admin/setting/system/index:get",
				"/admin/setting/system/getdir:get",
				"/admin/setting/system/view:get",
				"/admin/setting/system/index_redis:get",
				"/admin/setting/system/getdir_redis:get",
				"/admin/setting/system/view_redis:get",
				"article:get",
				"/admin/article/list:get",
				"/admin/article/save:post",
				"demo:get",
				"/admin/demo/show:get",
				"/admin/demo/upload:post",
			},
		},
		Cookies: suite.cookies,
	}
	// 发起post请求，以表单形式传递参数
	body, _ := httptestutil.PostForm("/admin/setting/admingroup/save", suite.router, param)
	assert.Contains(suite.T(), string(body), `"status":true`)
}

func TestExampleTestSuite(t *testing.T) {

	var err error

	err = configs.Init("")
	if err != nil {
		log.Fatalf("start fail:[Config Init] %s", err.Error())
	}

	err = web.Init()
	if err != nil {
		log.Fatalf("start fail:[Web Init] %s", err.Error())
	}

	err = redisx.Init()
	if err != nil {
		log.Fatalf("start fail:[Redis Init] %s", err.Error())
	}

	err = mysqlx.Init()
	if err != nil {
		log.Fatalf("start fail:[Mysql Init] %s", err.Error())
	}

	suite.Run(t, new(AdminTestSuite))
}
