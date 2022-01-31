/*
 * @Description:公用工具类
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */
package comment

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"math"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
获取项目根目录
*/
func RootPath() (path string, err error) {
	path = getCurrentAbPathByExecutable()
	if strings.Contains(path, getTmpDir()) {
		path = getCurrentAbPathByCaller()
	}
	path = strings.Replace(path, "/pkg/comment", "", 1)
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
生成随机字符串
*/
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

/**
密码加密
*/
func Encryption(password string, salt string) string {
	str := fmt.Sprintf("%s%s", password, salt)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
模板分页器
*/
type PageData struct {
	Count     int
	Data      interface{}
	Page      int
	PageHtml  template.HTML
	PageCount int
}

func PageOperation(c *gin.Context, db *gorm.DB, limit int, data interface{}) PageData {
	var count int64
	p := c.DefaultQuery("p", "1")

	page, _ := strconv.Atoi(p)

	db.Count(&count)

	db.Offset((page - 1) * limit).Limit(limit).Find(data)

	pageCount := int(math.Ceil(float64(count) / float64(limit)))

	url := c.FullPath()

	paramUrl := ""
	for k, v := range c.Request.URL.Query() {
		if k != "p" {
			paramUrl += "&" + k + "=" + v[0]
		}
	}

	pageHtml := "<nav aria-label='Page navigation'><ul class='pagination'><li><a href='" + url + "?p=1' aria-label='Previous'><span aria-hidden='true'>&laquo;</span></a></li>"

	if pageCount < 5 {
		for i := 1; i <= pageCount; i++ {
			if page == i {
				pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			} else {
				pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			}
		}
	} else {
		if page > 2 && page < pageCount-2 {
			for i := page - 2; i <= page+2; i++ {
				if page == i {
					pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				}
			}
		}

		if page <= 2 {
			var maxPage int
			if pageCount > 5 {
				maxPage = 5
			} else {
				maxPage = pageCount
			}
			for i := 1; i <= maxPage; i++ {

				if page == i {
					pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				}

			}
		}

		if page >= pageCount-2 {
			for i := pageCount - 4; i <= pageCount; i++ {

				if page == i {
					pageHtml += "<li class='active'><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
				}

			}
		}
	}
	pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(pageCount) + paramUrl + "' aria-label='Next'><span aria-hidden='true'>&raquo;</span></a></li><li></li></ul></nav>"

	if pageCount == 0 {
		pageHtml = ""
	}

	return PageData{
		Count:     1,
		Data:      data,
		Page:      1,
		PageHtml:  template.HTML(pageHtml),
		PageCount: pageCount,
	}
}

/**
*首字母大写
**/
func StrFirstToUpper(str string) (string, string) {
	temp := strings.Split(str, "_")
	var upperStr string
	var firstStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				firstStr += string(vv[i])
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr, firstStr
}

/*
*比较第二个slice一第一个slice的区别
 */
func CompareSlice(first []string, second []string) (add []string, incre []string) {

	secondMap := make(map[string]struct{})

	for _, v := range second {
		secondMap[v] = struct{}{}
	}

	for _, v := range first {
		_, ok := secondMap[v]
		if !ok {
			incre = append(incre, v)
		} else {
			delete(secondMap, v)
		}
	}

	for k, _ := range secondMap {
		add = append(add, k)
	}

	return
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
* 组装字符串
 */
func JoinStr(items ...interface{}) string {
	fmt.Println(items...)
	fmt.Println(len(items))
	if len(items) == 0 {
		return ""
	}
	var builder strings.Builder
	for _, v := range items {
		builder.WriteString(v.(string))
	}
	return builder.String()
}
