/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 14:18:10
 */
package article

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/paginater"

	"github.com/gin-gonic/gin"
)

type articleController struct {
	admin.BaseController
}

var Arc = articleController{}

func (con articleController) Add(c *gin.Context) {
	var article models.Article
	c.HTML(http.StatusOK, "article/article_form.html", gin.H{
		"article": article,
	})
}

func (con articleController) Edit(c *gin.Context) {
	articel_id := c.Query("article_id")

	id, _ := strconv.Atoi(articel_id)

	article, err := services.ArticleService.GetArticle(uint(id))
	if err != nil {
		con.Error(c, err.Error())
	}

	c.HTML(http.StatusOK, "article/article_form.html", gin.H{
		"article": article,
		"url":     c.Request.RequestURI,
	})
}

func (con articleController) List(c *gin.Context) {
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

	adminDb := services.ArticleService.GetArticles(req)

	articleData := paginater.PageOperation(c, adminDb, 1, &articleList)

	c.HTML(http.StatusOK, "article/article_list.html", gin.H{
		"articleData": articleData,
		"created_at":  c.Query("created_at"),
		"title":       c.Query("title"),
	})
}

func (con articleController) Save(c *gin.Context) {
	var (
		req models.ArticleReq
		err error
	)

	con.FormBind(c, &req)
	err = services.ArticleService.SaveArticle(req)
	if err != nil {
		con.Error(c, err.Error())
	}

	con.Success(c, "/admin/article/list", "添加成功")
}

func (con articleController) Del(c *gin.Context) {
	id := c.Query("article_id")
	fmt.Println(id)
	articleId, _ := strconv.Atoi(id)

	err := services.ArticleService.DelArticle(articleId)
	if err != nil {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "", "删除成功")
	}
}
