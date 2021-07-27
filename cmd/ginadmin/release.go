// +build !release

package main

import (
	_ "ginadmin/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	release = false
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
