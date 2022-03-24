/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-18 18:30:28
 */
package uploader

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/gphper/ginadmin/configs"
)

type LocalStorage struct {
}

func (stor LocalStorage) Save(file *multipart.FileHeader, dst string) (string, error) {

	var (
		dstFull  string
		filePath string
	)

	filePath = dst + "/" + file.Filename
	dstFull = configs.RootPath + "/" + filePath

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
	return filePath, err
}
