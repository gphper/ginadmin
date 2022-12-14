/**
 * @Author: GPHPER
 * @Date: 2022-12-14 10:07:09
 * @Github: https://github.com/gphper
 * @LastEditTime: 2022-12-14 10:27:39
 * @FilePath: \ginadmin\cmd\cli\file\model_gen_mod.go
 * @Description:
 */
package file

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gphper/ginadmin/configs"
	"github.com/spf13/cobra"
)

var modelStr string = `package models

import (
	"gorm.io/gorm"
)

type {{.Name}} struct {
	BaseModle
}

// 设置数据表名称
func ({{.Short}} *{{.Name}}) TableName() string {
	return "{{.Table}}"
}

// 填充数据
func ({{.Short}} *{{.Name}}) FillData(db *gorm.DB) {
	
}

// 设置数据库链接名称
func ({{.Short}} *{{.Name}}) GetConnName() string {
	return "default"
}
`

func writeModel(fileName string, firstName string) error {

	parms := struct {
		Name  string
		Table string
		Short string
	}{
		Name:  fileName,
		Table: modelName,
		Short: firstName,
	}

	newPath := configs.RootPath + "internal" + string(filepath.Separator) + "models" + string(filepath.Separator) + fileName + ".go"
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
