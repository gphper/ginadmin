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

var Uc = uploadController{}

func (con uploadController) Show(c *gin.Context) {

	c.HTML(http.StatusOK, "demo/upload.html", gin.H{})

}

func (con uploadController) Upload(c *gin.Context) {

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

	filepath, err := services.UpService.Save(stor, req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"path": filepath,
	})

}
