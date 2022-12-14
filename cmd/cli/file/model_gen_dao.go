/**
 * @Author: GPHPER
 * @Date: 2022-12-14 10:06:57
 * @Github: https://github.com/gphper
 * @LastEditTime: 2022-12-14 10:49:16
 * @FilePath: \ginadmin\cmd\cli\file\model_gen_dao.go
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

var daoStr string = `package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type {{.ModelName}}Dao struct {
	DB *gorm.DB
}

var (
	instance{{.ModelName}} *{{.ModelName}}Dao
	once{{.ModelName}}Dao  sync.Once
)

func New{{.ModelName}}Dao() *{{.ModelName}}Dao {
	once{{.ModelName}}Dao.Do(func() {
		instance{{.ModelName}} = &{{.ModelName}}Dao{DB: models.GetDB(&models.{{.ModelName}}{})}
	})
	return instance{{.ModelName}}
}

// 新增数据
func (dao *{{.ModelName}}Dao) Create(data models.{{.ModelName}}) error {
	return dao.DB.Create(&data).Error
}

// 获取单条数据
func (dao *{{.ModelName}}Dao) GetOne(conditions map[string]interface{}) (data models.{{.ModelName}}, err error) {

	err = dao.DB.First(&data, conditions).Error
	return
}

// 更新数据
func (dao *{{.ModelName}}Dao) UpdateColumns(conditions, field map[string]interface{}, tx *gorm.DB) error {

	if tx != nil {
		return tx.Model(&models.{{.ModelName}}{}).Where(conditions).UpdateColumns(field).Error
	}

	return dao.DB.Model(&models.{{.ModelName}}{}).Where(conditions).UpdateColumns(field).Error
}

// 删除数据
func (dao *{{.ModelName}}Dao) Del(conditions map[string]interface{}) error {
	return dao.DB.Delete(&models.{{.ModelName}}{}, conditions).Error
}
`

func writeDao(modelName string, file string) error {

	parms := struct {
		ModelName string // Person
	}{
		ModelName: modelName,
	}

	newDaoPath := configs.RootPath + "internal" + string(filepath.Separator) + "dao" + string(filepath.Separator) + file + "Dao.go"
	_, err := os.Lstat(newDaoPath)
	if err == nil {
		return errors.New("file already exist")
	}

	fileDao, err := os.Create(newDaoPath)
	if err != nil {
		cobra.CompError(err.Error())
		return err
	}
	defer fileDao.Close()

	tem, _ := template.New("models_file").Parse(daoStr)
	tem.ExecuteTemplate(fileDao, "models_file", parms)

	return nil
}
