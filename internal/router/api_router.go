/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-13 19:45:35
 */
package router

import (
	"github/gphper/ginadmin/internal/controllers/api/user"

	"github.com/gin-gonic/gin"
)

func ApiRouter(apiRouter *gin.RouterGroup) {
	apiRouter.Use()
	{
		apiUserRouter := apiRouter.Group("user")
		{
			apiUserRouter.POST("/example", user.Uc.UserExample)
		}
	}
}
