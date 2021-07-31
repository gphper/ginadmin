/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package main

import (
	"context"
	"ginadmin/configs"
	_ "ginadmin/internal/models"
	"ginadmin/internal/router"
	"ginadmin/pkg/comment"
	_ "ginadmin/pkg/cron"
	"ginadmin/web"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	swagHandler gin.HandlerFunc
	release     bool = true
)

// @title GinAdmin Api
// @version 1.0
// @description GinAdmin 示例项目

// @contact.name gphper
// @contact.url https://github.com/gphper/ginadmin

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:20011
// @basepath /api
func main() {

	var (
		rootPath   string
		separator  string
		uploadPath string
		err        error
	)

	//判断是否编译线上版本
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	rootPath, _ = comment.RootPath()
	separator = comment.GetLine()
	uploadPath = rootPath + separator + "uploadfile"

	r := router.Init()
	//判断是否添加api文档
	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}

	r.HTMLRender = web.LoadTemplates()
	r.StaticFS("/statics", web.StaticsFs)

	_, err = os.Stat(uploadPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(uploadPath, os.ModeDir)
			if err != nil {
				panic(err.Error())
			}
		}
	}
	r.StaticFS("/uploadfile", http.Dir(uploadPath))
	// pprof路由
	//pprof.Register(r)
	srv := &http.Server{
		Addr:    configs.App.BaseConf.Host + ":" + configs.App.BaseConf.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
