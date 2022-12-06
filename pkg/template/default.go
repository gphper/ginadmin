/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 11:58:44
 */
package template

import (
	"html/template"
	"reflect"
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
		"isset": func(param interface{}) bool {
			value := reflect.ValueOf(param)
			switch value.Kind() {
			case reflect.String:
				return value.Len() != 0
			case reflect.Bool:
				return value.Bool()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return value.Int() != 0
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return value.Uint() != 0
			case reflect.Float32, reflect.Float64:
				return value.Float() != 0
			case reflect.Interface, reflect.Ptr:
				return !value.IsNil()
			}
			return !reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
		},
	}
}
