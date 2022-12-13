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

type ApiRouter struct {
	root *gin.RouterGroup
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}

func (ar ApiRouter) addRouter(con IApiController, router *gin.RouterGroup) {
	con.Routes(router)
}

func (ar ApiRouter) AddRouters() {
	{
		apiUserRouter := ar.root.Group("user")
		{
			ar.addRouter(user.NewUserController(), apiUserRouter)
		}
	}
}
