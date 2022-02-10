/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:29:28
 */
package admin

import (
	"strings"

	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type articleService struct{}

var ArticleService = articleService{}

func (ser *articleService) GetArticle(articleId uint) (article models.Article, err error) {
	article.ArticleId = articleId
	err = dao.ArticleDao.DB.First(&article).Error
	if err != nil {
		return models.Article{}, err
	}
	return
}

func (ser *articleService) GetArticles(req models.ArticleIndexReq) (db *gorm.DB) {

	db = dao.AuDao.DB.Table("article")

	if req.Title != "" {
		db = db.Where("title like ?", "%"+req.Title+"%")
	}

	if req.CreatedAt != "" {
		period := strings.Split(req.CreatedAt, " ~ ")
		start := period[0] + " 00:00:00"
		end := period[1] + " 23:59:59"
		db = db.Where("created_at > ? ", start).Where("created_at < ?", end)
	}

	return
}

//添加或保存文章信息
func (ser *articleService) SaveArticle(req models.ArticleReq) (err error) {

	var (
		article models.Article
	)

	if req.ArticleId > 0 {

		article.ArticleId = uint(req.ArticleId)

		dao.ArticleDao.DB.First(&article)

		article.Title = req.Title
		article.Desc = req.Desc
		article.Content = req.Content
		article.CoverImg = req.CoverImg

		err = dao.ArticleDao.DB.Save(&article).Error
		if err != nil {
			return
		}

	} else {
		article.Title = req.Title
		article.Content = req.Content
		article.CoverImg = req.CoverImg
		article.Desc = req.Desc
		err = dao.ArticleDao.DB.Save(&article).Error
		if err != nil {
			return
		}
	}

	return
}

func (ser *articleService) DelArticle(id int) (err error) {
	var article models.Article
	article.ArticleId = uint(id)
	err = dao.ArticleDao.DB.Delete(&article).Error
	return
}
