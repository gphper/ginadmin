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
	store := cookie.NewStore([]byte("secret11111"))
	adminRouter.Use(sessions.Sessions("mysession", store))
	{
		adminRouter.GET("/captcha", admin.Lc.Captcha)
		/*******登录路由**********/
		adminRouter.GET("/login", admin.Lc.Login)
		adminRouter.POST("/login", admin.Lc.Login)
		adminRouter.GET("/login_out", admin.Lc.LoginOut)
		adminRouter.POST("/login_out", admin.Lc.LoginOut)

		adminHomeRouter := adminRouter.Group("/home")
		adminHomeRouter.Use(middleware.AdminUserAuth())
		{
			adminHomeRouter.GET("/", admin.Hc.Home)
			adminHomeRouter.GET("/welcome", admin.Hc.Welcome)
			adminHomeRouter.GET("/edit_password", admin.Hc.EditPassword)
			adminHomeRouter.POST("/save_password", admin.Hc.SavePassword)
			adminHomeRouter.POST("/save_skin", admin.Hc.SaveSkin)
		}

		adminSettingRouter := adminRouter.Group("/setting")
		adminSettingRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminGroup := adminSettingRouter.Group("/admingroup")
			{
				adminGroup.GET("/index", setting.Agc.Index)
				adminGroup.GET("/add", setting.Agc.AddIndex)
				adminGroup.POST("/save", setting.Agc.Save)
				adminGroup.GET("/edit", setting.Agc.Edit)
				adminGroup.GET("/del", setting.Agc.Del)
			}

			adminUser := adminSettingRouter.Group("/adminuser")
			{
				adminUser.GET("/index", setting.Auc.Index)
				adminUser.GET("/add", setting.Auc.AddIndex)
				adminUser.POST("/save", setting.Auc.Save)
				adminUser.GET("/edit", setting.Auc.Edit)
				adminUser.GET("/del", setting.Auc.Del)
			}

			adminSystem := adminSettingRouter.Group("/system")
			{
				adminSystem.GET("/index", setting.Asc.Index)
				adminSystem.GET("/getdir", setting.Asc.GetDir)
				adminSystem.GET("/view", setting.Asc.View)

				adminSystem.GET("/index_redis", setting.Asc.IndexRedis)
				adminSystem.GET("/getdir_redis", setting.Asc.GetDirRedis)
				adminSystem.GET("/view_redis", setting.Asc.ViewRedis)
			}

		}

		//Demo演示文件上传
		adminDemoRouter := adminRouter.Group("/demo")
		adminDemoRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminDemoRouter.GET("/show", demo.Uc.Show)
			adminDemoRouter.POST("/upload", demo.Uc.Upload)
		}

		//Article文章管理
		adminArticleRouter := adminRouter.Group("/article")
		adminArticleRouter.Use(middleware.AdminUserAuth(), middleware.AdminUserPrivs())
		{
			adminArticleRouter.GET("/list", article.Arc.List)
			adminArticleRouter.GET("/add", article.Arc.Add)
			adminArticleRouter.GET("/edit", article.Arc.Edit)
			adminArticleRouter.POST("/save", article.Arc.Save)
			adminArticleRouter.GET("/del", article.Arc.Del)
		}

		//文件上传
		adminUploadRouter := adminRouter.Group("/upload")
		adminUploadRouter.Use(middleware.AdminUserAuth())
		{
			adminUploadRouter.GET("/upload_html/:type_name/:id/:type/:now_num", upload.Upc.UploadHtml)
			adminUploadRouter.POST("/upload", upload.Upc.Upload)
		}

	}
}
