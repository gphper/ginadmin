package template

import (
	"html/template"
	"time"
)

var GlobalTemplateFun template.FuncMap

func init(){
	GlobalTemplateFun = template.FuncMap{
		"formatAsDate": func(t time.Time,format string)string {
			return t.Format(format)
		},
		"judgeContainPriv": func(privMap map[string]interface{},priv string)bool {
			//判断权限是all的全通过
			_,o :=privMap["all"]
			if o {
				return true
			}
			_,ok := privMap[priv]
			return ok
		},
		"pageOperate": func() template.HTML{
			return "<nav aria-label='Page navigation'><ul class='pagination'><li><a href='#' aria-label='Previous'><span aria-hidden='true'>&laquo;</span></a></li><li><a href='#'>1</a></li><li><a href='#'>2</a></li><li><a href='#' aria-label='Next'><span aria-hidden='true'>&raquo;</span></a></li></ul></nav>"
		},
	}
}
