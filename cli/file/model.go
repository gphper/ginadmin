package file

import (
	"bufio"
	"errors"
	"ginadmin/comment"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Params struct {
	Name  string
	Table string
	Short string
}

var modelStr string = `package models

type {{.Name}} struct {
	BaseModle
}

func ({{.Short}} *{{.Name}}) TableName() string {
	return "{{.Table}}"
}

func ({{.Short}} *{{.Name}}) FillData() {
	
}
`

var cmdModel = &cobra.Command{
	Use:   "model [-m modelName]",
	Short: "create model",
	Run:   modelFunc,
}

var modelName string

func init() {
	cmdModel.Flags().StringVarP(&modelName, "model", "m", "", "input model name eg: shop_items")
}

func strFirstToUpper(str string) (string, string) {
	temp := strings.Split(str, "_")
	var upperStr string
	var firstStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				firstStr += string(vv[i])
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr, firstStr
}

func modelFunc(cmd *cobra.Command, args []string) {
	if len(modelName) == 0 {
		cmd.Help()
	}
	fileName, firstName := strFirstToUpper(modelName)
	err := writeModel(fileName, firstName)

	if err != nil {
		cobra.CompErrorln(err.Error())
		return
	}
	modifyDefault(fileName)
}

func writeModel(fileName string, firstName string) error {

	parms := Params{
		Name:  fileName,
		Table: modelName,
		Short: firstName,
	}

	path, _ := comment.RootPath()
	newPath := path + "\\models\\" + fileName + ".go"
	_, err := os.Lstat(newPath)
	if err == nil {
		return errors.New("file already exist")
	}

	file, err := os.Create(newPath)
	if err != nil {
		cobra.CompError(err.Error())
		return err
	}
	defer file.Close()

	tem, _ := template.New("models_file").Parse(modelStr)
	tem.ExecuteTemplate(file, "models_file", parms)
	return nil
}

func modifyDefault(fileName string) {
	path, _ := comment.RootPath()
	oldPath := path + "\\models\\default.go"
	newPath := path + "\\models\\default_tmp.go"
	file, err := os.Open(oldPath)
	if err != nil {
		cobra.CompError(err.Error())
	}

	reader := bufio.NewReader(file)

	file_tmp, _ := os.Create(newPath)
	var flagTag int
	for {
		bytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if strings.Contains(string(bytes), "GetModels") {
			flagTag++
		}

		if flagTag > 0 {
			flagTag++
		}

		if flagTag == 4 {
			file_tmp.Write([]byte("\t\t&" + fileName + "{},\n"))
		}

		file_tmp.Write(bytes)
	}
	file.Close()
	file_tmp.Close()

	err = os.Remove(oldPath)
	if err != nil {
		cobra.CompError(err.Error())
	}
	err = os.Rename(newPath, oldPath)
	if err != nil {
		cobra.CompError(err.Error())
	}
}
