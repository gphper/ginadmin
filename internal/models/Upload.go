package models

import "mime/multipart"

type UploadReq struct {
	File *multipart.FileHeader `form:"file" label:"文件" binding:"required"`
	Dst  string                `form:"dst"`
}
