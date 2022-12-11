/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-13 19:45:35
 */
package router

import (
	"github.com/gphper/ginadmin/internal/controllers/api/user"

	"github.com/gin-gonic/gin"
)

func ApiRouter(apiRouter *gin.RouterGroup) {

	{
		apiUserRouter := apiRouter.Group("user")
		{
			addApiController(user.NewUserController(), apiUserRouter)

		}
	}
}
