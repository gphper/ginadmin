package template

import (
	"html/template"
	"time"
)

var GlobalTemplateFun template.FuncMap

func init() {
	GlobalTemplateFun = template.FuncMap{
		"formatAsDate": func(t time.Time, format string) string {
			return t.Format(format)
		},
		"judgeContainPriv": func(privMap map[string]struct{}, priv string) bool {
			//判断权限是all的全通过
			_, o := privMap["all"]
			if o {
				return true
			}
			_, ok := privMap[priv]
			return ok
		},
	}
}
