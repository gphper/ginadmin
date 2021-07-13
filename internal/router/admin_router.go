package router

import (
	"ginadmin/internal/controllers"
	"ginadmin/internal/controllers/demo"
	"ginadmin/internal/controllers/setting"
	"ginadmin/internal/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AdminRouter(adminRouter *gin.RouterGroup) {
	loginController := new(controllers.LoginController)
	homeController := new(controllers.HomeController)
	adminGroupController := new(setting.AdminGroupController)
	adminUserController := new(setting.AdminUserController)
	adminSysController := new(setting.AdminSystemController)
	uploadController := new(demo.UploadController)

	//设置后台用户权限中间件
	store := cookie.NewStore([]byte("secret11111"))
	adminRouter.Use(sessions.Sessions("mysession", store))
	{
		/*******登录路由**********/
		adminRouter.GET("/login", loginController.Login())
		adminRouter.POST("/login", loginController.Login())
		adminRouter.GET("/login_out", loginController.LoginOut())
		adminRouter.POST("/login_out", loginController.LoginOut())

		adminHomeRouter := adminRouter.Group("/home")
		adminHomeRouter.Use(middleware.AdminUserAuth())
		{
			adminHomeRouter.GET("/", homeController.Home())
			adminHomeRouter.GET("/welcome", homeController.Welcome())
			adminHomeRouter.GET("/edit_password", homeController.EditPassword())
			adminHomeRouter.POST("/save_password", homeController.SavePassword())
		}

		adminSettingRouter := adminRouter.Group("/setting")
		adminSettingRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminGroup := adminSettingRouter.Group("/admingroup")
			{
				adminGroup.GET("/index", adminGroupController.Index())
				adminGroup.GET("/add", adminGroupController.AddIndex())
				adminGroup.POST("/save", adminGroupController.Save())
				adminGroup.GET("/edit", adminGroupController.Edit())
				adminGroup.GET("/del", adminGroupController.Del())
			}

			adminUser := adminSettingRouter.Group("/adminuser")
			{
				adminUser.GET("/index", adminUserController.Index())
				adminUser.GET("/add", adminUserController.AddIndex())
				adminUser.POST("/save", adminUserController.Save())
				adminUser.GET("/edit", adminUserController.Edit())
				adminUser.GET("/del", adminUserController.Del())
			}

			adminSystem := adminSettingRouter.Group("/system")
			{
				adminSystem.GET("/index", adminSysController.Index())
				adminSystem.GET("/getdir", adminSysController.GetDir())
				adminSystem.GET("/view", adminSysController.View())
			}

		}

		//Demo演示文件上传
		adminDemoRouter := adminRouter.Group("/demo")
		adminDemoRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminDemoRouter.GET("/show", uploadController.Show())
			adminDemoRouter.POST("/upload", uploadController.Upload())
		}

	}
}
