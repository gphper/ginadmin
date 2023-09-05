package router

import "github.com/gin-gonic/gin"

type IController interface {
	Routes(*gin.RouterGroup)
}
