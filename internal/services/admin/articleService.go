/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-14 13:29:28
 */
package admin

import (
	"sync"

	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type articleService struct {
	Dao *dao.ArticleDao
}

var (
	instanceArticleService *articleService
	onceArticleService     sync.Once
)

func NewArticleService() *articleService {
	onceArticleService.Do(func() {
		instanceArticleService = &articleService{
			Dao: dao.NewArticleDao(),
		}
	})
	return instanceArticleService
}

func (ser *articleService) GetArticle(condition map[string]interface{}) (article models.Article, err error) {

	return ser.Dao.GetArticle(condition)
}

func (ser *articleService) GetArticles(req models.ArticleIndexReq) (db *gorm.DB) {

	return ser.Dao.GetArticles(req.Title, req.CreatedAt)
}

//添加或保存文章信息
func (ser *articleService) SaveArticle(req models.ArticleReq) (err error) {

	var (
		article models.Article
	)

	if req.ArticleId > 0 {

		err = ser.Dao.UpdateColumns(map[string]interface{}{
			"article_id": req.ArticleId,
		}, map[string]interface{}{
			"title":     req.Title,
			"desc":      req.Desc,
			"content":   req.Content,
			"cover_img": req.CoverImg,
		}, nil)
		if err != nil {
			return
		}

	} else {
		article.Title = req.Title
		article.Content = req.Content
		article.CoverImg = req.CoverImg
		article.Desc = req.Desc
		err = ser.Dao.DB.Save(&article).Error
		if err != nil {
			return
		}
	}

	return
}

func (ser *articleService) DelArticle(condition map[string]interface{}) (err error) {

	return ser.Dao.Del(condition)
}
