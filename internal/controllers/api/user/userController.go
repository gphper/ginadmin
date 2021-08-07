/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-04-18 19:07:39
 */
package user

import (
	"github/gphper/ginadmin/internal/controllers/api"
	"github/gphper/ginadmin/internal/models"
	"net/http"

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
// @Success 200 {object} models.UserReq
// @Router /user/example [post]
func (apicon *userController) UserExample(c *gin.Context) {
	var userReq models.UserReq
	c.ShouldBind(&userReq)
	c.JSON(http.StatusOK, userReq)
}
