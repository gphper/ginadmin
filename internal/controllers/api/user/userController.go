/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-04-18 19:07:39
 */
package user

import (
	"github/gphper/ginadmin/internal/controllers/api"
	"github/gphper/ginadmin/internal/models"
	apiservice "github/gphper/ginadmin/internal/services/api"

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
// @Param authorization header string true "token" default(Bearer)
// @Param info formData models.UserReq true "User info"
// @Success 200 {object} api.SuccessResponse{data=models.UserReq}
// @response default {object} api.DefaultResponse
// @Router /user/example [post]
func (apicon *userController) UserExample(c *gin.Context) {

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
// @Tags 用户注册
// @version 1.0
// @Accept multipart/form-data
// @Produce json
// @Param info formData models.UserRegisterReq true "User info"
// @Success 200 {object} api.SuccessResponse
// @response default {object} api.DefaultResponse
// @Router /user/register [post]
func (apicon *userController) Register(c *gin.Context) {
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
// @Tags 用户登录
// @version 1.0
// @Accept multipart/form-data
// @Produce json
// @Param info formData models.UserLoginReq true "Login info"
// @Success 200 {object} models.UserLoginRes
// @response default {object} api.DefaultResponse
// @Router /user/login [post]
func (apicon *userController) Login(c *gin.Context) {
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

func (apicon *userController) RefreshToken(c *gin.Context) {

}
