/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 14:18:10
 */
package article

import (
	"net/http"

	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/paginater"

	"github.com/gin-gonic/gin"
)

type articleController struct {
	admin.BaseController
}

func NewArticleController() articleController {
	return articleController{}
}

func (con articleController) Routes(rg *gin.RouterGroup) {
	rg.GET("/list", con.list)
	rg.GET("/add", con.add)
	rg.GET("/edit", con.edit)
	rg.POST("/save", con.save)
	rg.GET("/del", con.del)
}

func (con articleController) add(c *gin.Context) {
	c.HTML(http.StatusOK, "article/article_form.html", nil)
}

func (con articleController) edit(c *gin.Context) {

	articelId := c.Query("article_id")
	article, err := services.NewArticleService().GetArticle(map[string]interface{}{"article_id": articelId})
	if err != nil {
		con.Error(c, err.Error())
	}

	c.HTML(http.StatusOK, "article/article_form.html", gin.H{
		"article": article,
		"url":     c.Request.RequestURI,
	})
}

func (con articleController) list(c *gin.Context) {
	var (
		err         error
		req         models.ArticleIndexReq
		articleList []models.Article
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	adminDb := services.NewArticleService().GetArticles(req)

	articleData, err := paginater.PageOperation(c, adminDb, 1, &articleList)
	if err != nil {
		con.ErrorHtml(c, err)
		return
	}
	c.HTML(http.StatusOK, "article/article_list.html", gin.H{
		"articleData": articleData,
		"created_at":  c.Query("created_at"),
		"title":       c.Query("title"),
	})
}

func (con articleController) save(c *gin.Context) {
	var (
		req models.ArticleReq
		err error
	)

	con.FormBind(c, &req)
	err = services.NewArticleService().SaveArticle(req)
	if err != nil {
		con.Error(c, err.Error())
	}

	con.Success(c, "/admin/article/list", "添加成功")
}

func (con articleController) del(c *gin.Context) {

	id := c.Query("article_id")
	err := services.NewArticleService().DelArticle(map[string]interface{}{"article_id": id})
	if err != nil {
		con.Error(c, "删除失败")
		return
	}

	con.Success(c, "", "删除成功")
}
