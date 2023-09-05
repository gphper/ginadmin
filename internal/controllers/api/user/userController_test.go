/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-04-03 10:16:52
 */
package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/url"
	"testing"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/mysqlx"
	"github.com/gphper/ginadmin/pkg/redisx"
	"github.com/gphper/ginadmin/pkg/utils/httptestutil"
	"github.com/gphper/ginadmin/pkg/utils/httptestutil/router"
	"github.com/gphper/ginadmin/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// api UserController 相关接口测试
type ApiTestSuite struct {
	suite.Suite
	router *router.Router
	jtoken string
}

func (suite *ApiTestSuite) SetupSuite() {

	router, err := router.Init()
	if err != nil {
		log.Fatalf("router init fail: %s", err)
	}

	router.SetApiRoute("/api/user", NewUserController())

	suite.router = router
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
	assert.JSONEq(suite.T(), `{"code":200,"msg":"success","data":{}}`, string(body))
}

func (suite *ApiTestSuite) TestARegisterErr() {

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
	assert.JSONEq(suite.T(), `{"code":2000, "msg":"该邮箱已存在"}`, string(body))
}

func (suite *ApiTestSuite) TestARegisterBindErr() {

	option := httptestutil.OptionValue{
		Param: url.Values{
			"nickname": {"gphper"},
		},
	}
	// 发起post请求，以表单形式传递参数
	body, _ := httptestutil.PostForm("/api/user/register", suite.router, option)
	assert.JSONEq(suite.T(), `{"code":2001, "msg":"绑定参数出错"}`, string(body))
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

// 善后恢复数据处理
func (s *ApiTestSuite) TearDownSuite() {
	dao.NewUserDao().DB.Model(models.User{}).Where("email = ?", "570165887@qq.com").Delete(nil)
}

func TestApiSuite(t *testing.T) {

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

	suite.Run(t, new(ApiTestSuite))
}
