/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-20 21:46:59
 */
package dao

import (
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type uploadTypeDao struct {
	DB *gorm.DB
}

var UtDao = uploadTypeDao{DB: models.Db}
