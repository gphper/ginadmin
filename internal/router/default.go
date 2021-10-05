/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package router

import (
	"path/filepath"
	"time"

	"github/gphper/ginadmin/internal/middleware"
	"github/gphper/ginadmin/pkg/comment"
	"github/gphper/ginadmin/pkg/loggers/facade"
	"github/gphper/ginadmin/pkg/loggers/medium"
	"github/gphper/ginadmin/web"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	SwagHandler gin.HandlerFunc
)

func Init() *gin.Engine {

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/", "/metrics"})))
	router.Use(middleware.Metrics())
	prep(router)

	router.Use(medium.GinLog(facade.NewZaplog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewZaplog("admin"), true))
	router.Use(middleware.NotHttpStatusOk())
	// router.Use(medium.GinLog(facade.NewRedisLog("admin"), time.RFC3339, true), medium.RecoveryWithLog(facade.NewRedisLog("admin"), true))
	/*****admin路由定义******/
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	adminRouter := router.Group("/admin")

	AdminRouter(adminRouter)

	/***api路由定义****/
	apiRouter := router.Group("/api")

	ApiRouter(apiRouter)

	pprof.Register(router)

	return router
}

func prep(router *gin.Engine) {
	var (
		rootPath   string
		uploadPath string
		err        error
	)

	rootPath, _ = comment.RootPath()

	uploadPath = rootPath + string(filepath.Separator) + "uploadfile"

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
				panic(err.Error())
			}
		}
	}

	router.StaticFS("/uploadfile", http.Dir(uploadPath))
}
