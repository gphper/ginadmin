/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-15 20:14:09
 */
package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminUserDao struct {
	DB *gorm.DB
}

var insAud *adminUserDao
var onceAud sync.Once

func NewAdminUserDao() *adminUserDao {
	onceAud.Do(func() {
		insAud = &adminUserDao{DB: models.Db}
	})
	return insAud
}
