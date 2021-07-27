/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package router

import (
	"ginadmin/pkg/loggers"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.Default()

	router.Use(ginzap.Ginzap(loggers.AdminLogger, time.RFC3339, true), ginzap.RecoveryWithZap(loggers.AdminLogger, true))
	/*****admin路由定义******/
	adminRouter := router.Group("/admin")
	AdminRouter(adminRouter)

	/***api路由定义****/
	apiRouter := router.Group("/api")
	ApiRouter(apiRouter)
	return router
}
