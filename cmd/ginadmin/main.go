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
	template2 "ginadmin/pkg/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
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
// @host localhost:20010
// @basepath /api
func main() {

	//判断是否编译线上版本
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	rootPath, _ := comment.RootPath()
	separator := comment.GetLine()
	r := router.Init()

	//判断是否添加api文档
	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}

	r.HTMLRender = loadTemplates(rootPath + "/web/views")
	r.StaticFS("/statics", http.Dir(rootPath+separator+"web"+separator+"statics"))
	r.StaticFS("/uploadfile", http.Dir(rootPath+separator+"uploadfile"))
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

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layout/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/template/*/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		dirSlice := strings.Split(include, comment.GetLine())
		fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")
		r.AddFromFilesFuncs(fileName, template2.GlobalTemplateFun, files...)
	}
	return r
}
