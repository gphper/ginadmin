/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package router

import (
	"log"
	"path/filepath"
	"time"

	"net/http"
	"os"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/controllers"
	"github.com/gphper/ginadmin/internal/middleware"
	"github.com/gphper/ginadmin/pkg/loggers/facade"
	"github.com/gphper/ginadmin/pkg/loggers/medium"
	"github.com/gphper/ginadmin/pkg/utils/strings"
	"github.com/gphper/ginadmin/web"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	SwagHandler gin.HandlerFunc
)

func Init() *gin.Engine {

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.NoRoute(controllers.NewHandleController().Handle)
	router.NoMethod(controllers.NewHandleController().Handle)
	//初始化路由相关信息
	err := prep(router)
	if err != nil {
		log.Fatalf("start error route: %s", err.Error())
	}
	// router.Use(medium.GinLog(facade.NewZaplog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewZaplog("admin"), true))
	// router.Use(medium.GinLog(facade.NewRedisLog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewRedisLog("admin"), true))
	router.Use(middleware.Trace())
	router.Use(medium.GinLog(facade.NewLogger("admin"), time.RFC3339, false), medium.RecoveryWithLog(facade.NewLogger("admin"), true))
	/*****admin路由定义******/
	adminRouter := router.Group("/admin")

	AdminRouter(adminRouter)

	/***api路由定义****/
	apiRouter := router.Group("/api")

	ApiRouter(apiRouter)

	return router
}

func prep(router *gin.Engine) error {
	var (
		uploadPath string
		err        error
	)

	uploadPath = strings.JoinStr(configs.RootPath, string(filepath.Separator), "uploadfile")
	if SwagHandler != nil {
		router.GET("/swagger/*any", SwagHandler)
	}

	router.HTMLRender = web.LoadTemplates()

	router.StaticFS("/statics", web.StaticsFs)

	_, err = os.Stat(uploadPath)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(uploadPath, os.ModeDir)
			if err != nil {

				return err
			}
		}
	}

	router.StaticFS("/uploadfile", http.Dir(uploadPath))

	return nil
}
