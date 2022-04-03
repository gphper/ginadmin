/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-04-03 10:16:52
 */
package test

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gphper/ginadmin/internal/router"
	"github.com/gphper/ginadmin/pkg/utils/httptestutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ApiTestSuite struct {
	suite.Suite
	router *gin.Engine
	jtoken string
}

func (suite *ApiTestSuite) SetupSuite() {
	suite.router = router.Init()
}

// 注册接口测试
func (suite *ApiTestSuite) TestARegister() {

	option := httptestutil.OptionValue{
		Param: url.Values{
			"nickname":         {"gphper"},
			"password":         {"111111"},
			"confirm_password": {"111111"},
			"email":            {"570165887@qq.com"},
		},
	}
	// 发起post请求，以表单形式传递参数
	body, _ := httptestutil.PostForm("/api/user/register", suite.router, option)
	assert.JSONEq(suite.T(), `{"code":1,"msg":"success","data":{}}`, string(body))

}

// 登录接口测试
func (suite *ApiTestSuite) TestBLogin() {

	type (
		Data struct {
			Jtoken  string
			Retoken string
		}

		response struct {
			Code int
			Msg  string
			Data Data
		}
	)

	var resp response

	option := httptestutil.OptionValue{
		Param: url.Values{
			"password": {"111111"},
			"email":    {"570165887@qq.com"},
		},
	}

	body, _ := httptestutil.PostForm("/api/user/login", suite.router, option)
	assert.Contains(suite.T(), string(body), "jtoken")

	decode := json.NewDecoder(bytes.NewReader(body))
	if err := decode.Decode(&resp); err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	suite.jtoken = resp.Data.Jtoken
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}
