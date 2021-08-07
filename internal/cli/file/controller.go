package file

import (
	"fmt"
	"github/gphper/ginadmin/pkg/comment"
	"html/template"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type CParams struct {
	ClassName string
	Pagename  string
	Methods   []string
}

var controllerStr string = `package {{.Pagename}}

import (
	"github/gphper/ginadmin/controllers"
	"github.com/gin-gonic/gin"
)

type {{.ClassName}} struct {
	controllers.BaseController
}

{{range .Methods}}
func (con *{{$.ClassName}}) {{.}}() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}
{{end}}
`

var cmdController = &cobra.Command{
	Use:   "controller [-p pagename -c controllerName -l methods]",
	Short: "create controller file",
	Run:   controllerFunc,
}

var (
	pagename       string
	controllerName string
	methods        string
)

func init() {
	cmdController.Flags().StringVarP(&pagename, "pagename", "p", "", "input pagename eg: setting")
	cmdController.Flags().StringVarP(&controllerName, "controllerName", "c", "", "input controller name eg: AdminController")
	cmdController.Flags().StringVarP(&methods, "methods", "m", "", "input methods eg: index,add,del")
}

func controllerFunc(cmd *cobra.Command, args []string) {
	if len(pagename) == 0 || len(controllerName) == 0 {
		cmd.Help()
		return
	}
	err := writeController()
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		return
	}
}

func writeController() error {
	parms := CParams{
		ClassName: controllerName,
		Pagename:  pagename,
	}

	mSlice := strings.Split(methods, ",")
	if len(mSlice) != 0 {
		parms.Methods = mSlice
	}

	path, _ := comment.RootPath()

	basePath := path + "\\controllers\\" + pagename
	_, err := os.Lstat(basePath)
	if err != nil {
		os.Mkdir(basePath, os.ModeDir)
	}

	newPath := basePath + "\\" + controllerName + ".go"
	_, err = os.Lstat(newPath)
	if err == nil {
		return fmt.Errorf("%s file already exist", controllerName+".go")
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
