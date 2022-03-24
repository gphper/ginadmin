/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-20 20:03:25
 */
package upload

import (
	"net/http"
	"strings"

	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	adminServ "github.com/gphper/ginadmin/internal/services/admin"
	"github.com/gphper/ginadmin/pkg/uploader"

	"github.com/gin-gonic/gin"
)

type uploadController struct {
	admin.BaseController
}

var Upc = uploadController{}

func (con uploadController) UploadHtml(c *gin.Context) {

	var (
		req         models.UploadHtmlReq
		uploadType  models.UploadType
		err         error
		mimeMap     map[string]string
		mimeTypes   string
		extensions  string
		maxAllowNum uint
	)

	err = con.UriBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	mimeMap = make(map[string]string, 11)
	mimeMap["gif"] = "image/gif"
	mimeMap["jpg"] = "image/jpeg"
	mimeMap["jpeg"] = "image/jpeg"
	mimeMap["png"] = "image/png"
	mimeMap["xml"] = "text/xml"
	mimeMap["svg"] = "image/svg+xml"
	mimeMap["xls"] = "application/vnd.ms-excel"
	mimeMap["xlsx"] = "application/vnd.ms-excel"
	mimeMap["zip"] = "application/zip"
	mimeMap["rar"] = "application/x-rar-compressed"
	mimeMap["mp3"] = "audio/mpeg"

	err = dao.UtDao.DB.Where("type_name", req.TypeName).First(&uploadType).Error
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	types := strings.Split(uploadType.AllowType, "|")
	slices := make([]string, len(types))
	for typek, typev := range types {
		if mimeType, ok := mimeMap[typev]; ok {
			slices[typek] = mimeType
		}
	}
	mimeTypes = strings.Join(slices, ",")
	oldAllowNum := uploadType.AllowNum

	if uploadType.AllowNum > 1 {
		uploadType.AllowNum = uploadType.AllowNum - req.NowNum
	}

	sizeMB := uploadType.AllowSize / 1024 / 1024

	extensions = strings.ReplaceAll(uploadType.AllowType, "|", ",")
	maxAllowNum = uploadType.AllowNum - 1

	c.HTML(http.StatusOK, "upload/upload.html", gin.H{
		"mimeTypes":     mimeTypes,
		"info":          uploadType,
		"extensions":    extensions,
		"file_id":       req.Id,
		"type":          req.Type,
		"now_num":       req.NowNum,
		"old_allow_num": oldAllowNum,
		"sizeMB":        sizeMB,
		"maxAllowNum":   maxAllowNum,
	})
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

	req.Dst = "uploadfile"
	filepath, err := adminServ.UpService.Save(stor, req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file_path": filepath,
	})

}
