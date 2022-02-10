/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-24 19:42:52
 */
package internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gphper/ginadmin/configs"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Route *gin.Engine
}

func (app Application) Run() {

	srv := &http.Server{
		Addr:    configs.App.BaseConf.Host + ":" + configs.App.BaseConf.Port,
		Handler: app.Route,
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
