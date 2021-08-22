/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package router

import (
	"time"

	"github/gphper/ginadmin/pkg/loggers/facade"
	"github/gphper/ginadmin/pkg/loggers/medium"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.Default()

	// router.Use(medium.GinLog(facade.NewZaplog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewZaplog("admin"), true))
	router.Use(medium.GinLog(facade.NewRedisLog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewRedisLog("admin"), true))
	/*****admin路由定义******/
	adminRouter := router.Group("/admin")
	AdminRouter(adminRouter)

	/***api路由定义****/
	apiRouter := router.Group("/api")
	ApiRouter(apiRouter)
	return router
}
