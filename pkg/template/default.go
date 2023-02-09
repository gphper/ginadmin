/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 11:58:44
 */
package template

import (
	"html/template"
	"strings"
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
			ok, err := casbinauth.Check(username, obj, act)
			if !ok || err != nil {
				return false
			}
			return true
		},
		"judegContainSlicePriv": func(username string, objs []string) bool {

			for _, obj := range objs {
				priv := strings.Split(obj, ":")
				ok, err := casbinauth.Check(username, priv[0], priv[1])
				if !ok || err != nil {
					return false
				}
			}

			return true
		},
		"joinSlicePriv": func(objs []string) string {
			return strings.Join(objs, "|")
		},
		"judegInMap": func(find string, items map[string]struct{}) bool {
			_, ok := items[find]
			return ok
		},
	}
}
