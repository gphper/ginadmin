/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-20 21:46:59
 */
package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/mysqlx"

	"gorm.io/gorm"
)

type uploadTypeDao struct {
	DB *gorm.DB
}

var insUtd *uploadTypeDao
var onceUtd sync.Once

func NewUploadTypeDao() *uploadTypeDao {
	onceUtd.Do(func() {
		insUtd = &uploadTypeDao{DB: mysqlx.GetDB(&models.UploadType{})}
	})
	return insUtd
}
