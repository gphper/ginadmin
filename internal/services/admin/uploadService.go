/*
 * @Description:上传附件服务
 * @Author: gphper
 * @Date: 2021-07-18 17:52:20
 */
package admin

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/uploader"
)

type uploadService struct {
}

var (
	instanceUploadService *uploadService
	onceUploadService     sync.Once
)

func NewUploadService() *uploadService {
	onceUploadService.Do(func() {
		instanceUploadService = &uploadService{}
	})
	return instanceUploadService
}

func (ser *uploadService) Save(storage uploader.Storage, req models.UploadReq) (string, error) {
	return storage.Save(req.File, req.Dst)
}
