/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 11:58:44
 */
package template

import (
	"html/template"
	"time"

	"github.com/gphper/ginadmin/pkg/casbinauth"
)

var GlobalTemplateFun template.FuncMap

func init() {
	GlobalTemplateFun = template.FuncMap{
		"formatAsDate": func(t time.Time, format string) string {
			return t.Format(format)
		},
		"judgeContainPriv": func(username string, obj string, act string) bool {
			if username == "admin" {
				return true
			}
			ok, err := casbinauth.Check(username, obj, act)
			if !ok || err != nil {
				return false
			}
			return true
		},
		"judegInMap": func(find string, items map[string]struct{}) bool {
			_, ok := items[find]
			return ok
		},
		"isset": func(param string) bool {
			return param != ""
		},
	}
}
