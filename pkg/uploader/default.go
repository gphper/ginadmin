package uploader

import "mime/multipart"

type Storage interface {
	Save(file *multipart.FileHeader, dst string) (string, error)
}
