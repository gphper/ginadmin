package demo

import (
	"ginadmin/internal/controllers/admin"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type uploadController struct {
	admin.BaseController
}

var Uc = uploadController{}

func (con *uploadController) Show() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo/upload.html", gin.H{})
	}
}

func (con *uploadController) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 单文件
		file, _ := c.FormFile("upload")
		log.Println(file.Filename)
		filepath := "uploadfile/" + file.Filename
		// 上传文件至指定目录
		c.SaveUploadedFile(file, filepath)

		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		c.JSON(http.StatusOK, gin.H{
			"path": filepath,
		})
	}
}
