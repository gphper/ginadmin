/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-15 20:14:09
 */
package dao

import (
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminUserDao struct {
	DB *gorm.DB
}

var AuDao = adminUserDao{DB: models.Db}
