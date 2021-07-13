package router

import (
	"ginadmin/api/apiuser"

	"github.com/gin-gonic/gin"
)

func ApiRouter(apiRouter *gin.RouterGroup) {
	apiUser := new(apiuser.ApiUserController)
	apiRouter.Use()
	{
		apiUserRouter := apiRouter.Group("user")
		{
			apiUserRouter.GET("/list", apiUser.UserList())
		}
	}
}
