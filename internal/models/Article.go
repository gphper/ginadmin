/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 14:17:14
 */
package models

import "time"

type Article struct {
	BaseModle
	ArticleId uint   `gorm:"primary_key;auto_increment"`
	Title     string `gorm:"size:100;comment:'标题'"`
	Desc      string `gorm:"size:100;comment:'描述'"`
	CoverImg  string `gorm:"size:100;comment:'封面图'"`
	Content   string `gorm:"type:text;comment:'文章内容'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleIndexReq struct {
	Title     string `form:"title"`
	CreatedAt string `form:"created_at"`
}

type ArticleReq struct {
	ArticleId int    `form:"article_id"`
	Title     string `form:"title" label:"标题" binding:"required"`
	CoverImg  string `form:"cover_img" label:"封面图" binding:"required"`
	Content   string `form:"content" label:"文章详情" binding:"required"`
	Desc      string `form:"desc" label:"文章描述" binding:"required"`
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) FillData() {

}
