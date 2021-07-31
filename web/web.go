//+build !embed

/*
 * @Description:部署拷贝静态文件时忽略该文件
 * @Author: gphper
 * @Date: 2021-07-31 10:59:00
 */
package web

import (
	"ginadmin/pkg/comment"
	"net/http"
	"path/filepath"
	"strings"

	template2 "ginadmin/pkg/template"

	"github.com/gphper/multitemplate"
)

var rootPath string
var StaticsFs http.FileSystem

func init() {
	rootPath, _ = comment.RootPath()
	separator := comment.GetLine()
	StaticsFs = http.Dir(rootPath + separator + "web" + separator + "statics")
}

func LoadTemplates() multitemplate.Renderer {
	templatesDir := rootPath + "/web/views"
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layout/*.html")

	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/template/*/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		dirSlice := strings.Split(include, comment.GetLine())
		fileName := strings.Join(dirSlice[len(dirSlice)-2:], "/")
		r.AddFromFilesFuncs(fileName, template2.GlobalTemplateFun, files...)
	}
	return r
}
