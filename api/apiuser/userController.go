package apiuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiUserController struct {
}

func (apicon *ApiUserController) UserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello world",
		})
	}
}
