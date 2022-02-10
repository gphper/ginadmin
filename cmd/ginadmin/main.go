/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gphper/ginadmin/internal/router"
	_ "github.com/gphper/ginadmin/pkg/cron"

	"github.com/gphper/ginadmin/internal"

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
	showLogo()

	//判断是否编译线上版本
	if release {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	app := internal.Application{
		Route: router.Init(),
	}
	app.Run()

}

func showLogo() {
	fmt.Println("   _____ _                   _           _       ")
	fmt.Println("  / ____(_)         /\\      | |         (_)      ")
	fmt.Println(" | |  __ _ _ __    /  \\   __| |_ __ ___  _ _ __  ")
	fmt.Println(" | | |_ | | '_ \\  / /\\ \\ / _` | '_ ` _ \\| | '_ \\ ")
	fmt.Println(" | |__| | | | | |/ _____\\ (_| | | | | | | | | | |")
	fmt.Println("  \\_____|_|_| |_/_/    \\_\\__,_|_| |_| |_|_|_| |_| \n")
}
