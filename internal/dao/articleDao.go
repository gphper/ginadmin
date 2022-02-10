/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:39:32
 */
package dao

import (
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type articleDao struct {
	DB *gorm.DB
}

var ArticleDao = articleDao{DB: models.Db}
