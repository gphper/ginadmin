/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-27 11:02:04
 */
package paginater

import (
	"html/template"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
模板分页器
*/
type PageData struct {
	Count     int
	Data      interface{}
	Page      int
	PageHtml  template.HTML
	PageCount int
}

func PageOperation(c *gin.Context, db *gorm.DB, limit int, data interface{}) PageData {
	var count int64
	p := c.DefaultQuery("p", "1")

	page, _ := strconv.Atoi(p)

	db.Count(&count)

	db.Offset((page - 1) * limit).Limit(limit).Find(data)

	pageCount := int(math.Ceil(float64(count) / float64(limit)))

	url := c.FullPath()

	paramUrl := ""
	for k, v := range c.Request.URL.Query() {
		if k != "p" {
			paramUrl += "&" + k + "=" + v[0]
		}
	}

	pageHtml := "<nav aria-label='Page navigation'><ul class='pagination'><li><a href='" + url + "?p=1' aria-label='Previous'><span aria-hidden='true'>&laquo;</span></a></li>"

	if pageCount < 5 {
		for i := 1; i <= pageCount; i++ {
			if page == i {
				pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			} else {
				pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			}
		}
	} else {
		if page > 2 && page < pageCount-2 {
			for i := page - 2; i <= page+2; i++ {
				if page == i {
					pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				}
			}
		}

		if page <= 2 {
			var maxPage int
			if pageCount > 5 {
				maxPage = 5
			} else {
				maxPage = pageCount
			}
			for i := 1; i <= maxPage; i++ {

				if page == i {
					pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				}

			}
		}

		if page >= pageCount-2 {
			for i := pageCount - 4; i <= pageCount; i++ {

				if page == i {
					pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				}

			}
		}
	}
	pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(pageCount) + paramUrl + "' aria-label='Next'><span aria-hidden='true'>&raquo;</span></a></li><li></li></ul></nav>"

	if pageCount == 0 {
		pageHtml = ""
	}

	return PageData{
		Count:     1,
		Data:      data,
		Page:      1,
		PageHtml:  template.HTML(pageHtml),
		PageCount: pageCount,
	}
}
