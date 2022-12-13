package router

import "github.com/gin-gonic/gin"

// 后台控制器接口
type IAdminController interface {
	Success(*gin.Context, string, string)
	Error(*gin.Context, string)
	ErrorHtml(*gin.Context, error)
	Html(*gin.Context, int, string, gin.H)
	FormBind(*gin.Context, interface{}) error
	UriBind(*gin.Context, interface{}) error
	Routes(*gin.RouterGroup)
}

// API控制器接口
type IApiController interface {
	Success(*gin.Context, interface{})
	Error(*gin.Context, error)
	FormBind(*gin.Context, interface{}) error
	Routes(*gin.RouterGroup)
}
