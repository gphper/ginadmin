//go:build embed
// +build embed

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

func Init() error {
	static, err := fs.Sub(StaticPath, "statics")
	if err != nil {
		return err
	}
	StaticsFs = http.FS(static)

	return nil
}

func LoadTemplates() (render multitemplate.Renderer, err error) {

	render = multitemplate.NewRenderer()

	layouts, err := fs.Glob(viewPath, "views/layout/*.html")
	if err != nil {
		return
	}

	includes, err := fs.Glob(viewPath, "views/template/*/*.html")
	if err != nil {
		return
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		dirSlice := strings.Split(include, "/")
		fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")
		render.AddFromFsFuncs(fileName, template2.GlobalTemplateFun, viewPath, files...)
	}
	return
}
