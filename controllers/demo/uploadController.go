package demo

import (
	"ginadmin/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UploadController struct {
	controllers.BaseController
}

func (con *UploadController) Show() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo/upload.html", gin.H{})
	}
}

func (con *UploadController) Upload() gin.HandlerFunc {
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
