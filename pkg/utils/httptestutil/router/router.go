/**
 * @Author: GPHPER
 * @Date: 2022-12-12 15:06:04
 * @Github: https://github.com/gphper
 * @LastEditTime: 2022-12-13 19:51:24
 * @FilePath: \ginadmin\pkg\router\router.go
 * @Description:
 */
package router

import (
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gphper/ginadmin/internal"
	"github.com/gphper/multitemplate"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(r *gin.Engine) Router {
	return Router{
		r: r,
	}
}

func (route Router) SetGlobalMiddleware(middlewares ...gin.HandlerFunc) {
	route.r.Use(middlewares...)
}

// 设置自定义模板加载
func (route Router) SetHtmlRenderer(render multitemplate.Renderer) {
	route.r.HTMLRender = render
}

// 设置swagger访问
func (route Router) SetSwaagerHandle(path string, handle gin.HandlerFunc) {
	route.r.GET(path, handle)
}

// 设置静态路径
func (route Router) SetStaticFile(path string, fs http.FileSystem) {
	route.r.StaticFS(path, fs)
}

// 设置附件保存地址
func (route Router) SetUploadDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModeDir)
			if err != nil {

				return err
			}
		}
	}

	route.r.StaticFS("/uploadfile", http.Dir(path))

	return nil
}

func (route Router) SetEngine(app *internal.Application) {
	app.Route = route.r
}

func (route Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route.r.ServeHTTP(w, req)
}

func (route Router) SetRouteError(handler gin.HandlerFunc) {
	route.r.NoMethod(handler)
	route.r.NoRoute(handler)
}

func (route Router) SetApiRoute(path string, ic IController, middlewares ...gin.HandlerFunc) {
	ar := route.r.Group(path)
	if len(middlewares) > 0 {
		ar.Use(middlewares...)
	}
	ic.Routes(ar)
}

func (route Router) SetAdminRoute(path string, ic IController, middlewares ...gin.HandlerFunc) {

	store := cookie.NewStore([]byte("1GdFRMs4fcWBvLXT"))
	middlewares = append(middlewares, gzip.Gzip(gzip.DefaultCompression), sessions.Sessions("mysession", store))

	ar := route.r.Group(path)
	if len(middlewares) > 0 {
		ar.Use(middlewares...)
	}
	ic.Routes(ar)
}
