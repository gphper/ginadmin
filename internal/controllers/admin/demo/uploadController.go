/*
 * @Description:上传附件示例
 * @Author: gphper
 * @Date: 2021-04-18 19:07:39
 */

package demo

import (
	"net/http"

	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/models"
	services "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/uploader"

	"github.com/gin-gonic/gin"
)

type uploadController struct {
	admin.BaseController
}

func NewUploadController() uploadController {
	return uploadController{}
}

func (con uploadController) Routes(rg *gin.RouterGroup) {
	rg.GET("/show", con.show)
	rg.POST("/upload", con.upload)
}

func (con uploadController) show(c *gin.Context) {

	con.Html(c, http.StatusOK, "demo/upload.html", gin.H{})

}

func (con uploadController) upload(c *gin.Context) {

	var (
		err error
		req models.UploadReq
	)
	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	req.Dst = "uploadfile"

	stor := uploader.LocalStorage{}

	filepath, err := services.NewUploadService().Save(stor, req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"path": filepath,
	})

}
