/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:39:32
 */
package dao

import (
	"strings"
	"sync"

	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/mysqlx"

	"gorm.io/gorm"
)

type ArticleDao struct {
	DB *gorm.DB
}

var (
	instanceArticle *ArticleDao
	onceArticleDao  sync.Once
)

func NewArticleDao() *ArticleDao {
	onceArticleDao.Do(func() {
		instanceArticle = &ArticleDao{DB: mysqlx.GetDB(&models.Article{})}
	})
	return instanceArticle
}

func (dao *ArticleDao) GetArticle(conditions map[string]interface{}) (article models.Article, err error) {

	err = dao.DB.First(&article, conditions).Error
	return
}

func (dao *ArticleDao) GetArticles(title string, createdAt string) (db *gorm.DB) {

	db = dao.DB.Table("article")

	if title != "" {
		db = db.Where("title like ?", "%"+title+"%")
	}

	if createdAt != "" {
		period := strings.Split(createdAt, " ~ ")
		start := period[0] + " 00:00:00"
		end := period[1] + " 23:59:59"
		db = db.Where("created_at > ? ", start).Where("created_at < ?", end)
	}

	return
}

func (dao *ArticleDao) UpdateColumns(conditions, field map[string]interface{}, tx *gorm.DB) error {

	if tx != nil {
		return tx.Model(&models.Article{}).Where(conditions).UpdateColumns(field).Error
	}

	return dao.DB.Model(&models.Article{}).Where(conditions).UpdateColumns(field).Error
}

func (dao *ArticleDao) Del(conditions map[string]interface{}) error {
	return dao.DB.Delete(&models.Article{}, conditions).Error
}
