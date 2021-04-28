package main

import (
	"context"
	"ginadmin/comment"
	_ "ginadmin/comment/cron"
	template2 "ginadmin/comment/template"
	"ginadmin/conf"
	_ "ginadmin/models"
	"ginadmin/router"
	"github.com/gin-contrib/multitemplate"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

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
		fileName := strings.Join(dirSlice[2:], "/")
		r.AddFromFilesFuncs(fileName, template2.GlobalTemplateFun, files...)
	}
	return r
}

func main() {
	r := router.Init()
	r.HTMLRender = loadTemplates("./views")
	r.StaticFS("/statics", http.Dir("./statics"))
	r.StaticFS("/uploadfile", http.Dir("./uploadfile"))
	// pprof路由
	//pprof.Register(r)
	srv := &http.Server{
		Addr:    conf.App.BaseConf.Host + ":" + conf.App.BaseConf.Port,
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
