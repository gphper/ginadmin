/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:39:32
 */
package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type articleDao struct {
	DB *gorm.DB
}

var insAd *articleDao
var onceAd sync.Once

func NewArticleDao() *articleDao {
	onceAd.Do(func() {
		insAd = &articleDao{DB: models.Db}
	})
	return insAd
}
