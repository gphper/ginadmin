/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-13 19:45:35
 */
package router

import (
	"github.com/gphper/ginadmin/internal/controllers/api/user"
	"github.com/gphper/ginadmin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouter(apiRouter *gin.RouterGroup) {

	{
		apiUserRouter := apiRouter.Group("user")
		{
			apiUserRouter.POST("/example", user.Uc.UserExample)
			apiUserRouter.POST("/register", user.Uc.Register)
			apiUserRouter.POST("/login", user.Uc.Login)
			apiUserRouter.POST("/refresh", user.Uc.RefreshToken)
		}

		apiExampleRouter := apiRouter.Group("example")
		apiExampleRouter.Use(middleware.JwtAuth())
		{
			apiExampleRouter.POST("/index", user.Uc.UserExample)
		}
	}
}
