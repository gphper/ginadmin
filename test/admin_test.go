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
}

func (suite *AdminTestSuite) TestLoginGet() {

	body := httptestutil.Get("/admin/login", suite.router)

	assert.Contains(suite.T(), string(body), "登录")
}

func (suite *AdminTestSuite) TestLoginPost() {
	param := url.Values{
		"username": {"admin"},
		"password": {"111111"},
	}
	// 发起post请求，以表单形式传递参数
	body, cookies := httptestutil.PostForm("/admin/login", param, suite.router)

	if len(cookies) != 0 {
		suite.cookies = cookies
	}

	assert.Contains(suite.T(), string(body), `"status":true`)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(AdminTestSuite))
}
