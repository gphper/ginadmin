package run

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gphper/ginadmin/configs"
	_ "github.com/gphper/ginadmin/docs"
	"github.com/gphper/ginadmin/internal"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/internal/redis"
	"github.com/gphper/ginadmin/internal/router"
	"github.com/gphper/ginadmin/web"
	"github.com/spf13/cobra"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var CmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run app",
	Run:   runFunction,
}

var (
	configPath string
	mode       string
)

func init() {
	CmdRun.Flags().StringVarP(&configPath, "config path", "c", "", "config path")
	CmdRun.Flags().StringVarP(&mode, "mode", "m", "dev", "dev or release")
}

func runFunction(cmd *cobra.Command, args []string) {
	showLogo()

	configs.Init(configPath)

	web.Init()

	redis.Init()
	models.Init()

	showPanel()
	//判断是否编译线上版本
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	} else {
		router.SwagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
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

func showPanel() {
	fmt.Println("App running at:")
	fmt.Printf("- Http Address:   %c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 34, "http://"+configs.App.Base.Host+":"+configs.App.Base.Port, 0x1B)
	fmt.Println("")
}
