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
		admin.NewLoginController().Routes(adminRouter)

		adminHomeRouter := adminRouter.Group("/home")
		adminHomeRouter.Use(middleware.AdminUserAuth())
		{
			admin.NewHomeController().Routes(adminHomeRouter)
		}

		adminSettingRouter := adminRouter.Group("/setting")
		adminSettingRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminGroup := adminSettingRouter.Group("/admingroup")
			{
				setting.NewAdminGroupController().Routes(adminGroup)
			}

			adminUser := adminSettingRouter.Group("/adminuser")
			{
				setting.NewAdminUserController().Routes(adminUser)
			}

			adminSystem := adminSettingRouter.Group("/system")
			{
				setting.NewAdminSystemController().Routes(adminSystem)
			}

		}

		//Demo演示文件上传
		adminDemoRouter := adminRouter.Group("/demo")
		adminDemoRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			demo.NewUploadController().Routes(adminDemoRouter)
		}

		//Article文章管理
		adminArticleRouter := adminRouter.Group("/article")
		adminArticleRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			article.NewArticleController().Routes(adminArticleRouter)
		}

		//文件上传
		adminUploadRouter := adminRouter.Group("/upload")
		adminUploadRouter.Use(middleware.AdminUserAuth())
		{
			upload.NewUploadController().Routes(adminUploadRouter)
		}

	}
}
