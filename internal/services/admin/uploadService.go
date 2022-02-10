/*
 * @Description:上传附件服务
 * @Author: gphper
 * @Date: 2021-07-18 17:52:20
 */
package admin

import (
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/uploader"
)

type uploadService struct{}

var UpService = uploadService{}

func (ser *uploadService) Save(storage uploader.Storage, req models.UploadReq) (string, error) {
	return storage.Save(req.File, req.Dst)
}
