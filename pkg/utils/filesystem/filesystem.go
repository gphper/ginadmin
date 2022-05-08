/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-27 10:57:10
 */
package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

/**
获取项目根目录
*/
func RootPath() (path string, err error) {
	path = getCurrentAbPathByExecutable()
	if strings.Contains(path, getTmpDir()) {
		path = getCurrentAbPathByCaller()
	}
	path = strings.Replace(path, "pkg/utils/filesystem", "", 1)
	return
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	if dir == "" {
		dir = "tmp"
	}

	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

/**
* 打开文件句柄
**/
func OpenFile(filepath string) (file *os.File, err error) {

	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		return
	}

	dir := path.Dir(filepath)
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, fs.FileMode(os.O_CREATE))
			if err != nil {
				return
			}
		}
	}
	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	return
}

/**
* 过滤非法访问的路径
 */
func FilterPath(root, path string) (string, error) {

	newPath := fmt.Sprintf("%s%s", root, path)
	absPath, err := filepath.Abs(newPath)
	if err != nil {
		return "", err
	}

	absPath = filepath.FromSlash(absPath)
	ifOver := strings.HasPrefix(absPath, filepath.FromSlash(root))
	if !ifOver {
		return "", errors.New("access to the path is prohibited")
	}

	return absPath, nil
}
