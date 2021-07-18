package services

import (
	"ginadmin/internal/models"
	"ginadmin/pkg/uploader"
)

type uploadService struct{}

var UpService = uploadService{}

func (ser *uploadService) Save(storage uploader.Storage, req models.UploadReq) (string, error) {
	return storage.Save(req.File, req.Dst)
}
