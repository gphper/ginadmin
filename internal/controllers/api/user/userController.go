/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-04-18 19:07:39
 */
package user

import (
	"github.com/gphper/ginadmin/internal/controllers/api"
	"github.com/gphper/ginadmin/internal/models"
	apiservice "github.com/gphper/ginadmin/internal/services/api"

	"github.com/gin-gonic/gin"
)

type userController struct {
	api.BaseController
}

var Uc = userController{}

// @Summary 展示用户信息
// @Id 1
// @Tags 示例
// @version 1.0
// @Accept multipart/form-data
// @Produce json
// @Param authorization header string true "token"
// @Param info formData models.UserReq true "User info"
// @Success 200 {object} api.SuccessResponse{data=models.UserReq}
// @response default {object} api.DefaultResponse
// @Router /example/index [post]
func (apicon userController) UserExample(c *gin.Context) {

	var (
		err     error
		userReq models.UserReq
	)
	err = apicon.FormBind(c, &userReq)

	if err != nil {
		apicon.Error(c, err)
		return
	}

	apicon.Success(c, userReq)
}

// @Summary 用户注册
// @Id 2
// @Tags 用户
// @version 1.0
// @Accept multipart/form-data
// @Produce json
// @Param info formData models.UserRegisterReq true "User info"
// @Success 200 {object} api.SuccessResponse
// @response default {object} api.DefaultResponse
// @Router /user/register [post]
func (apicon userController) Register(c *gin.Context) {
	var (
		err error
		req models.UserRegisterReq
	)

	err = apicon.FormBind(c, &req)
	if err != nil {
		apicon.Error(c, err)
		return
	}

	err = apiservice.UserService.Register(req)
	if err != nil {
		apicon.Error(c, err)
		return
	}

	apicon.Success(c, gin.H{})
}

// @Summary 用户登录
// @Id 3
// @Tags 用户
// @version 1.0
// @Accept multipart/form-data
// @Produce json
// @Param info formData models.UserLoginReq true "Login info"
// @Success 200 {object} models.UserLoginRes
// @response default {object} api.DefaultResponse
// @Router /user/login [post]
func (apicon userController) Login(c *gin.Context) {
	var (
		err     error
		req     models.UserLoginReq
		jtoken  string
		retoken string
		resp    models.UserLoginRes
	)

	err = apicon.FormBind(c, &req)
	if err != nil {
		apicon.Error(c, err)
		return
	}

	jtoken, retoken, err = apiservice.UserService.Login(req)
	if err != nil {
		apicon.Error(c, err)
		return
	}

	resp.Jtoken = jtoken
	resp.Retoken = retoken

	apicon.Success(c, resp)
}

// @Summary 刷新jtoken
// @Id 4
// @Tags 用户
// @version 1.0
// @Accept multipart/form-data
// @Produce json
// @Param info formData models.UserRefreshTokenReq true "info"
// @Success 200 {json} {"code":1,"msg":"success","data":{"jtoken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHAiOiIyMDIxLTEyLTI2VDE5OjI1OjI4Ljg0OTIzNzUrMDg6MDAiLCJOYW1lIjoiZ3BocGVyIiwiVWlkIjo0fQ==.ab81bb7134978afe976df55b45789aefd10f6c3edb969bae283c32c080083b89"}}
// @response default {object} api.DefaultResponse
// @Router /user/refresh [post]
func (apicon userController) RefreshToken(c *gin.Context) {
	var (
		err    error
		req    models.UserRefreshTokenReq
		jtoken string
	)

	err = apicon.FormBind(c, &req)
	if err != nil {
		apicon.Error(c, err)
		return
	}

	jtoken, err = apiservice.UserService.RefreshToken(req)
	if err != nil {
		apicon.Error(c, err)
		return
	}

	apicon.Success(c, gin.H{
		"jtoken": jtoken,
	})
}
