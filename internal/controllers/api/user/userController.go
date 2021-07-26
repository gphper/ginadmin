/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-04-18 19:07:39
 */
package user

import (
	"ginadmin/internal/controllers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	api.BaseController
}

var Uc = userController{}

// @Summary 获取用户信息
// @Id 1
// @Tags 用户
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "Account ID"
// @Router /article [post]
func (apicon *userController) UserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
