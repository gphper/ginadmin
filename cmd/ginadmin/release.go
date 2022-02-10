// +build !release

package main

import (
	_ "github.com/gphper/ginadmin/docs"
	"github.com/gphper/ginadmin/internal/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	release = false
	router.SwagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
