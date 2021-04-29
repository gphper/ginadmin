package router

import (
	"ginadmin/api/apiuser"
	"ginadmin/controllers"
	"ginadmin/controllers/demo"
	"ginadmin/controllers/setting"
	"ginadmin/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.Default()
	store := cookie.NewStore([]byte("secret11111"))
	router.Use(sessions.Sessions("mysession", store))
	//router.Use(ginzap.Ginzap(loggers.AdminLogger, time.RFC3339, true))
	//router.Use(ginzap.RecoveryWithZap(loggers.AdminLogger, true))
	/*****admin路由定义******/
	adminRouter := router.Group("/admin")
	//设置后台用户权限中间件
	adminRouter.Use(sessions.Sessions("mysession", store))
	{
		/*******登录路由**********/
		adminRouter.GET("/login", (&controllers.LoginController{}).Login())
		adminRouter.POST("/login", (&controllers.LoginController{}).Login())
		adminRouter.GET("/login_out", (&controllers.LoginController{}).LoginOut())
		adminRouter.POST("/login_out", (&controllers.LoginController{}).LoginOut())

		adminHomeRouter := adminRouter.Group("/home")
		adminHomeRouter.Use(middleware.AdminUserAuth())
		{
			adminHomeRouter.GET("/", (&controllers.HomeController{}).Home())
			adminHomeRouter.GET("/welcome", (&controllers.HomeController{}).Welcome())
			adminHomeRouter.GET("/edit_password", (&controllers.HomeController{}).EditPassword())
			adminHomeRouter.POST("/save_password", (&controllers.HomeController{}).SavePassword())
		}

		adminSettingRouter := adminRouter.Group("/setting")
		adminSettingRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminGroup := adminSettingRouter.Group("/admingroup")
			{
				adminGroup.GET("/index", (&setting.AdminGroupController{}).Index())
				adminGroup.GET("/add", (&setting.AdminGroupController{}).AddIndex())
				adminGroup.POST("/save", (&setting.AdminGroupController{}).Save())
				adminGroup.GET("/edit", (&setting.AdminGroupController{}).Edit())
				adminGroup.GET("/del", (&setting.AdminGroupController{}).Del())
			}

			adminUser := adminSettingRouter.Group("/adminuser")
			{
				adminUser.GET("/index", (&setting.AdminUserController{}).Index())
				adminUser.GET("/add", (&setting.AdminUserController{}).AddIndex())
				adminUser.POST("/save", (&setting.AdminUserController{}).Save())
				adminUser.GET("/edit", (&setting.AdminUserController{}).Edit())
				adminUser.GET("/del", (&setting.AdminUserController{}).Del())
			}

			adminSystem := adminSettingRouter.Group("/system")
			{
				adminSystem.GET("/index", (&setting.AdminSystemController{}).Index())
				adminSystem.GET("/getdir", (&setting.AdminSystemController{}).GetDir())
				adminSystem.GET("/view", (&setting.AdminSystemController{}).View())
			}

		}

		//Demo演示文件上传
		adminDemoRouter := adminRouter.Group("/demo")
		adminDemoRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminDemoRouter.GET("/show", (&demo.UploadController{}).Show())
			adminDemoRouter.POST("/upload", (&demo.UploadController{}).Upload())
		}

	}

	/***api路由定义****/
	apiRouter := router.Group("/api")
	apiRouter.Use()
	{
		apiUserRouter := apiRouter.Group("user")
		{
			apiUserRouter.GET("/list", (&apiuser.ApiUserController{}).UserList())
		}
	}

	return router
}
