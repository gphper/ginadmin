/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-13 19:45:15
 */
package router

import (
	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/controllers/admin/article"
	"github.com/gphper/ginadmin/internal/controllers/admin/demo"
	"github.com/gphper/ginadmin/internal/controllers/admin/setting"
	"github.com/gphper/ginadmin/internal/controllers/admin/upload"
	"github.com/gphper/ginadmin/internal/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AdminRouter(adminRouter *gin.RouterGroup) {

	//设置后台用户权限中间件
	store := cookie.NewStore([]byte("1GdFRMs4fcWBvLXT"))
	adminRouter.Use(sessions.Sessions("mysession", store))
	{

		addAdminController(admin.NewLoginController(), adminRouter)

		adminHomeRouter := adminRouter.Group("/home")
		adminHomeRouter.Use(middleware.AdminUserAuth())
		{
			addAdminController(admin.NewHomeController(), adminHomeRouter)
		}

		adminSettingRouter := adminRouter.Group("/setting")
		adminSettingRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminGroup := adminSettingRouter.Group("/admingroup")
			{
				addAdminController(setting.NewAdminGroupController(), adminGroup)
			}

			adminUser := adminSettingRouter.Group("/adminuser")
			{
				addAdminController(setting.NewAdminUserController(), adminUser)
			}

			adminSystem := adminSettingRouter.Group("/system")
			{
				addAdminController(setting.NewAdminSystemController(), adminSystem)
			}

		}

		//Demo演示文件上传
		adminDemoRouter := adminRouter.Group("/demo")
		adminDemoRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			addAdminController(demo.NewUploadController(), adminDemoRouter)
		}

		//Article文章管理
		adminArticleRouter := adminRouter.Group("/article")
		adminArticleRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			addAdminController(article.NewArticleController(), adminArticleRouter)
		}

		//文件上传
		adminUploadRouter := adminRouter.Group("/upload")
		adminUploadRouter.Use(middleware.AdminUserAuth())
		{
			addAdminController(upload.NewUploadController(), adminUploadRouter)
		}

	}
}
