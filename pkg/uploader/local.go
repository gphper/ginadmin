package uploader

import (
	"io"
	"mime/multipart"
	"os"
)

type LocalStorage struct {
}

func (stor LocalStorage) Save(file *multipart.FileHeader, dst string) (string, error) {
	var dstFull = dst + "/" + file.Filename
	src, err := file.Open()
	if err != nil {
		return dstFull, err
	}
	defer src.Close()

	out, err := os.Create(dstFull)
	if err != nil {
		return dstFull, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return dstFull, err
}
