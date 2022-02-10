//+build embed

/*
 * @Description:部署拷贝静态文件时忽略该文件
 * @Author: gphper
 * @Date: 2021-07-31 10:28:43
 */

package web

import (
	"embed"
	"fmt"
	template2 "github.com/gphper/ginadmin/pkg/template"
	"github.com/gphper/multitemplate"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed statics
var StaticPath embed.FS

//go:embed views
var viewPath embed.FS

var StaticsFs http.FileSystem

func init() {
	static, err := fs.Sub(StaticPath, "statics")
	if err != nil {
		panic(err.Error())
	}
	StaticsFs = http.FS(static)
}

func LoadTemplates() multitemplate.Renderer {

	r := multitemplate.NewRenderer()

	layouts, err := fs.Glob(viewPath, "views/layout/*.html")

	fmt.Println(layouts)

	if err != nil {
		panic(err.Error())
	}
	includes, err := fs.Glob(viewPath, "views/template/*/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		dirSlice := strings.Split(include, "/")
		fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")
		r.AddFromFsFuncs(fileName, template2.GlobalTemplateFun, viewPath, files...)
	}
	return r
}
