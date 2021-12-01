/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-18 18:30:28
 */
package uploader

import (
	"github/gphper/ginadmin/pkg/comment"
	"io"
	"mime/multipart"
	"os"
)

type LocalStorage struct {
}

func (stor LocalStorage) Save(file *multipart.FileHeader, dst string) (string, error) {

	var (
		root     string
		dstFull  string
		filePath string
	)

	root, _ = comment.RootPath()
	filePath = dst + "/" + file.Filename
	dstFull = root + "/" + filePath

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
