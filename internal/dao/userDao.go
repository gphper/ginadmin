/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-22 09:52:35
 */
package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type userDao struct {
	DB *gorm.DB
}

var insUd *userDao
var onceUd sync.Once

func NewUserDao() *userDao {
	onceUd.Do(func() {
		insUd = &userDao{DB: models.Db}
	})
	return insUd
}
