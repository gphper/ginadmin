/*
 * @Description:系统管理
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */

package setting

import (
	"bufio"
	"context"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/controllers/admin"
	"github.com/gphper/ginadmin/pkg/loggers"
	"github.com/gphper/ginadmin/pkg/redisx"
	"github.com/gphper/ginadmin/pkg/utils/filesystem"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"

	"github.com/gin-gonic/gin"
)

type adminSystemController struct {
	admin.BaseController
}

func NewAdminSystemController() adminSystemController {
	return adminSystemController{}
}

func (con adminSystemController) Routes(rg *gin.RouterGroup) {
	rg.GET("/index", con.index)
	rg.GET("/getdir", con.getDir)
	rg.GET("/view", con.view)
	rg.GET("/index_redis", con.indexRedis)
	rg.GET("/getdir_redis", con.getDirRedis)
	rg.GET("/view_redis", con.viewRedis)
}

/**
日志目录页面
*/
func (con adminSystemController) index(c *gin.Context) {

	var (
		path     string
		err      error
		log_path string
	)

	path = gstrings.JoinStr(configs.RootPath, string(filepath.Separator), "logs")

	files, err := ioutil.ReadDir(path)

	log_path = gstrings.JoinStr(string(filepath.Separator), "logs")

	ctx, _ := c.Get("ctx")

	if err != nil {
		loggers.LogError(ctx.(context.Context), "admin", "读取目录失败", map[string]string{"error": err.Error()})
		con.ErrorHtml(c, err)
		return
	}

	con.Html(c, http.StatusOK, "setting/systemlog.html", gin.H{
		"log_path": log_path,
		"files":    files,
		"line":     string(filepath.Separator),
	})
}

/**
获取目录
*/
func (con adminSystemController) getDir(c *gin.Context) {

	type FileNode struct {
		Name string `json:"name"`
		Path string `json:"path"`
		Type string `json:"type"`
	}

	var (
		path      string
		err       error
		fileSlice []FileNode
		files     []fs.FileInfo
	)

	fileSlice = make([]FileNode, 0)
	path, err = filesystem.FilterPath(configs.RootPath+"logs", c.Query("path"))
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	files, err = ioutil.ReadDir(path)
	if err != nil {
		con.Error(c, "获取目录失败")
		return
	}

	for _, v := range files {
		var fileType string
		if v.IsDir() {
			fileType = "dir"
		} else {
			fileType = "file"
		}
		fileSlice = append(fileSlice, FileNode{
			Name: v.Name(),
			Path: gstrings.JoinStr(c.Query("path"), string(filepath.Separator), v.Name()),
			Type: fileType,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fileSlice,
	})
}

/**
获取日志详情
*/
func (con adminSystemController) view(c *gin.Context) {

	var (
		err       error
		startLine int
		endLine   int
		scanner   *bufio.Scanner
		line      int
	)

	startLine, err = strconv.Atoi(c.DefaultQuery("start_line", "1"))
	if err != nil {
		con.ErrorHtml(c, err)
		return
	}
	endLine, err = strconv.Atoi(c.DefaultQuery("end_line", "20"))
	if err != nil {
		con.ErrorHtml(c, err)
		return
	}

	var filecontents []string
	filePath, err := filesystem.FilterPath(configs.RootPath+"logs", c.Query("path"))
	if err != nil {
		con.ErrorHtml(c, err)
		return
	}

	fi, err := os.Open(filePath)
	if err != nil {
		con.ErrorHtml(c, err)
		return
	}
	defer fi.Close()

	scanner = bufio.NewScanner(fi)
	for scanner.Scan() {
		line++
		if line >= startLine && line <= endLine {
			// 在要求行数内取得数据
			filecontents = append(filecontents, scanner.Text())
		} else {
			continue
		}
	}

	con.Html(c, http.StatusOK, "setting/systemlog_view.html", gin.H{
		"file_path":    c.Query("path"),
		"filecontents": filecontents,
		"start_line":   startLine,
		"end_line":     endLine,
		"line":         line,
	})

}

/**
日志目录页面
*/
func (con adminSystemController) indexRedis(c *gin.Context) {

	path := "logs"

	dateSlice, err := redisx.GetRedisClient().Keys("logs:*").Result()

	ctx, _ := c.Get("ctx")

	if err != nil {
		loggers.LogError(ctx.(context.Context), "admin", "读取目录失败", map[string]string{"error": err.Error()})
		con.ErrorHtml(c, err)
		return
	}

	dates := make(map[string]struct{})

	for _, v := range dateSlice {
		temp := strings.Split(v, ":")

		if _, ok := dates[temp[1]]; !ok {
			dates[temp[1]] = struct{}{}
		}
	}

	con.Html(c, http.StatusOK, "setting/systemlog_redis.html", gin.H{
		"log_path": path,
		"files":    dates,
	})
}

/**
获取目录
*/
func (con adminSystemController) getDirRedis(c *gin.Context) {

	path := c.Query("path")

	type FileNode struct {
		Name string `json:"name"`
		Path string `json:"path"`
		Type string `json:"type"`
	}

	pathSlice := strings.Split(path, "_")

	pattern := pathSlice[0] + ":*"

	dateSlice, err := redisx.GetRedisClient().Keys(pattern).Result()

	ctx, _ := c.Get("ctx")

	if err != nil {
		loggers.LogError(ctx.(context.Context), "admin", "读取目录失败", map[string]string{"error": err.Error()})
		con.ErrorHtml(c, err)
		return
	}

	fileSlice := make([]FileNode, 0)

	tempMap := make(map[string]struct{})

	for _, v := range dateSlice {
		temp := strings.Split(v, ":")
		index, _ := strconv.Atoi(pathSlice[1])
		var fileType string

		if index+2 == len(temp) {
			fileType = "file"
		} else {
			fileType = "dir"
		}

		if _, ok := tempMap[temp[index+1]]; ok {
			continue
		} else {
			tempMap[temp[index+1]] = struct{}{}
		}

		fileSlice = append(fileSlice, FileNode{
			Name: temp[index+1],
			Path: pathSlice[0] + ":" + temp[index+1] + "_" + strconv.Itoa(index+1),
			Type: fileType,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"data": fileSlice,
	})
}

/**
获取日志详情
*/
func (con adminSystemController) viewRedis(c *gin.Context) {

	startLine, _ := strconv.Atoi(c.DefaultQuery("start_line", "1"))

	endLine, _ := strconv.Atoi(c.DefaultQuery("end_line", "20"))

	filePath := c.Query("path")

	pathSlice := strings.Split(filePath, "_")

	filecontents, _ := redisx.GetRedisClient().LRange(pathSlice[0], int64(startLine-1), int64(endLine-1)).Result()

	line, _ := redisx.GetRedisClient().LLen(pathSlice[0]).Result()

	con.Html(c, http.StatusOK, "setting/systemlog_viewredis.html", gin.H{
		"file_path":    filePath,
		"filecontents": filecontents,
		"start_line":   startLine,
		"end_line":     endLine,
		"line":         line,
	})

}
