/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-31 19:59:19
 */
package test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gphper/ginadmin/internal/router"
	"github.com/gphper/ginadmin/pkg/utils/httptestutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AdminTestSuite struct {
	suite.Suite
	router  *gin.Engine
	cookies []*http.Cookie
}

func (suite *AdminTestSuite) SetupSuite() {
	suite.router = router.Init()
	//先执行登录操作获取cookie
	suite.TestLoginPost()
}

// 登录页测试
func (suite *AdminTestSuite) TestLoginGet() {

	body := httptestutil.Get("/admin/login", suite.router)

	assert.Contains(suite.T(), string(body), "登录")
}

// 登录
func (suite *AdminTestSuite) TestLoginPost() {

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
func (suite *AdminTestSuite) TestAddGroup() {
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
	suite.Run(t, new(AdminTestSuite))
}
