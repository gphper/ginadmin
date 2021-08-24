/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package main

import (
	"github/gphper/ginadmin/internal/router"
	_ "github/gphper/ginadmin/pkg/cron"

	"github/gphper/ginadmin/internal"

	"github.com/gin-gonic/gin"
)

var (
	release bool = true
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

	//判断是否编译线上版本
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	app := internal.Application{
		Route: router.Init(),
	}
	app.Run()

}
