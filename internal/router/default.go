/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package router

import (
	"path/filepath"
	"time"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/controllers"
	"github.com/gphper/ginadmin/internal/middleware"
	"github.com/gphper/ginadmin/pkg/loggers/facade"
	"github.com/gphper/ginadmin/pkg/loggers/medium"
	"github.com/gphper/ginadmin/pkg/utils/strings"
	"github.com/gphper/ginadmin/web"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() (*Router, error) {

	router := NewRouter(gin.Default())

	//设置404错误处理
	router.SetRouteError(controllers.NewHandleController().Handle)

	//设置全局中间件
	router.SetGlobalMiddleware(middleware.Trace(), medium.GinLog(facade.NewLogger("admin"), time.RFC3339, false), medium.RecoveryWithLog(facade.NewLogger("admin"), true))

	// 开发模式设置接口文档路由
	if gin.Mode() == "debug" {
		router.SetSwaagerHandle("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 设置模板解析函数
	render, err := web.LoadTemplates()
	if err != nil {

		return nil, err
	}
	router.SetHtmlRenderer(render)

	//设置静态资源
	router.SetStaticFile("/statics", web.StaticsFs)

	//设置上传附件
	uploadPath := strings.JoinStr(configs.RootPath, string(filepath.Separator), "uploadfile")
	err = router.SetUploadDir(uploadPath)
	if err != nil {
		return nil, err
	}

	// 设置后台全局中间件
	store := cookie.NewStore([]byte("1GdFRMs4fcWBvLXT"))
	router.SetAdminRoute(NewAdminRouter(), gzip.Gzip(gzip.DefaultCompression), sessions.Sessions("mysession", store))
	router.SetApiRoute(NewApiRouter())
	return &router, nil
}
