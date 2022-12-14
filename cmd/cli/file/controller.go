// go run .\cmd\ginadmin\ file controller -p=shop -c=shopController -t=admin
package file

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gphper/ginadmin/configs"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"

	"github.com/spf13/cobra"
)

var controllerStr string = `package {{.PackageName}}

import (
	"github.com/gphper/ginadmin/internal/controllers/{{.TypeName}}"

	"github.com/gin-gonic/gin"
)

type {{.ClassName}} struct {
	{{.TypeName}}.BaseController
}

func New{{.UpClassName}}() {{.ClassName}}  {
	return {{.ClassName}}{}
}

func (con {{.ClassName}}) Routes(rg *gin.RouterGroup) {
	{{range $kk,$vv := .Methods}}
	rg.{{$vv}}("/{{$kk}}", con.{{$kk}})
	{{end}}
}

{{range $k,$v := .Methods}}
func (con {{$.ClassName}}) {{$k}}(c *gin.Context) {
}
{{end}}
`

var cmdController = &cobra.Command{
	Use:   "controller [-p pagename -c controllerName -m methods]",
	Short: "create controller file",
	Run:   controllerFunc,
}

var (
	pagename       string
	controllerName string
	methods        string
	typename       string
)

func init() {
	cmdController.Flags().StringVarP(&typename, "typename", "t", "", "input typename api or admin")
	cmdController.Flags().StringVarP(&pagename, "pagename", "p", "", "input pagename eg: setting")
	cmdController.Flags().StringVarP(&controllerName, "controllerName", "c", "", "input controller name eg: AdminController")
	cmdController.Flags().StringVarP(&methods, "methods", "m", "list:get,add:get,save:post,edit:get,del:get", "input methods eg: index:get,add:get")
	cmdController.MarkFlagRequired("typename")
}

func controllerFunc(cmd *cobra.Command, args []string) {

	//判断 typename 类型
	if typename != "admin" && typename != "api" {
		fmt.Println("typename: api or admin")
		os.Exit(1)
	}

	err := configs.Init("")
	if err != nil {

		fmt.Printf("start fail:[Config Init] %s", err.Error())
		os.Exit(1)
	}

	if len(pagename) == 0 || len(controllerName) == 0 {
		cmd.Help()
		return
	}

	pageSlice := strings.Split(pagename, "\\")
	packageName := pageSlice[len(pageSlice)-1]

	upName, _, _ := gstrings.StrFirstToUpper(packageName)

	err = writeController(upName, packageName)
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		os.Exit(1)
	}
}

func writeController(upName string, packageName string) error {
	parms := struct {
		ClassName   string
		Pagename    string
		PackageName string
		UpClassName string
		TypeName    string
		Methods     map[string]string
	}{
		ClassName:   controllerName,
		Pagename:    pagename,
		PackageName: packageName,
		TypeName:    typename,
		UpClassName: upName,
	}

	methods := strings.Split(methods, ",")
	methodMap := make(map[string]string)
	for _, v := range methods {
		methodSlice := strings.Split(v, ":")
		methodMap[methodSlice[0]] = strings.ToUpper(methodSlice[1])
	}
	parms.Methods = methodMap

	basePath := configs.RootPath + "internal" + string(filepath.Separator) + "controllers" + string(filepath.Separator) + typename + string(filepath.Separator) + pagename
	_, err := os.Lstat(basePath)
	if err != nil {
		os.Mkdir(basePath, os.ModeDir)
	}

	newPath := basePath + string(filepath.Separator) + controllerName + ".go"
	_, err = os.Lstat(newPath)
	if err == nil {
		return err
	}

	file, err := os.Create(newPath)
	if err != nil {
		cobra.CompError(err.Error())
		return err
	}
	defer file.Close()

	tem, _ := template.New("controller_file").Parse(controllerStr)
	tem.ExecuteTemplate(file, "controller_file", parms)
	return nil
}
