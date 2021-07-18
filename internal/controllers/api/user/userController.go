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

func (apicon *userController) UserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello world",
		})
	}
}
