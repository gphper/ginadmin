package user

import (
	"ginadmin/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	controllers.BaseController
}

func(con *UserController) Index() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"hello world",
		})
	}
}
