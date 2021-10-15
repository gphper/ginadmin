/*
 * @Description:系统管理
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */

package setting

import (
	"bufio"
	"fmt"
	"github/gphper/ginadmin/internal/controllers/admin"
	"github/gphper/ginadmin/internal/redis"
	"github/gphper/ginadmin/pkg/comment"
	"github/gphper/ginadmin/pkg/loggers"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type adminSystemController struct {
	admin.BaseController
}

var Asc = adminSystemController{}

/**
日志目录页面
*/
func (con *adminSystemController) Index(c *gin.Context) {

	path, err := comment.RootPath()
	if err != nil {
		fmt.Printf("get root path err:%v", err)
	}
	path = path + string(filepath.Separator) + "logs"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		loggers.LogError("admin", "读取目录失败", map[string]string{"error": err.Error()})
	}
	c.HTML(http.StatusOK, "setting/systemlog.html", gin.H{
		"log_path": path,
		"files":    files,
		"line":     string(filepath.Separator),
	})
}

/**
获取目录
*/
func (con *adminSystemController) GetDir(c *gin.Context) {

	type FileNode struct {
		Name string `json:"name"`
		Path string `json:"path"`
		Type string `json:"type"`
	}
	fileSlice := make([]FileNode, 0)
	path := c.Query("path")
	files, err := ioutil.ReadDir(path)
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
			Path: path + string(filepath.Separator) + v.Name(),
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
func (con *adminSystemController) View(c *gin.Context) {

	startLine, _ := strconv.Atoi(c.DefaultQuery("start_line", "1"))
	endLine, _ := strconv.Atoi(c.DefaultQuery("end_line", "20"))
	var filecontents []string
	filePath := c.Query("path")
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	line := 0
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line++
		if line >= startLine && line <= endLine {
			// 在要求行数内取得数据
			filecontents = append(filecontents, scanner.Text())
		} else {
			continue
		}
	}

	c.HTML(http.StatusOK, "setting/systemlog_view.html", gin.H{
		"file_path":    filePath,
		"filecontents": filecontents,
		"start_line":   startLine,
		"end_line":     endLine,
		"line":         line,
	})

}

/**
日志目录页面
*/
func (con *adminSystemController) IndexRedis(c *gin.Context) {

	path := "logs"

	dateSlice, err := redis.RedisClient.Keys("logs:*").Result()

	if err != nil {
		loggers.LogError("admin", "读取目录失败", map[string]string{"error": err.Error()})
	}

	dates := make(map[string]struct{})

	for _, v := range dateSlice {
		temp := strings.Split(v, ":")

		if _, ok := dates[temp[1]]; !ok {
			dates[temp[1]] = struct{}{}
		}
	}

	c.HTML(http.StatusOK, "setting/systemlog_redis.html", gin.H{
		"log_path": path,
		"files":    dates,
	})
}

/**
获取目录
*/
func (con *adminSystemController) GetDirRedis(c *gin.Context) {

	path := c.Query("path")

	type FileNode struct {
		Name string `json:"name"`
		Path string `json:"path"`
		Type string `json:"type"`
	}

	pathSlice := strings.Split(path, "_")

	pattern := pathSlice[0] + ":*"

	dateSlice, err := redis.RedisClient.Keys(pattern).Result()

	if err != nil {
		loggers.LogError("admin", "读取目录失败", map[string]string{"error": err.Error()})
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
func (con *adminSystemController) ViewRedis(c *gin.Context) {

	startLine, _ := strconv.Atoi(c.DefaultQuery("start_line", "1"))

	endLine, _ := strconv.Atoi(c.DefaultQuery("end_line", "20"))

	filePath := c.Query("path")

	pathSlice := strings.Split(filePath, "_")

	filecontents, _ := redis.RedisClient.LRange(pathSlice[0], int64(startLine-1), int64(endLine-1)).Result()

	line, _ := redis.RedisClient.LLen(pathSlice[0]).Result()

	c.HTML(http.StatusOK, "setting/systemlog_viewredis.html", gin.H{
		"file_path":    filePath,
		"filecontents": filecontents,
		"start_line":   startLine,
		"end_line":     endLine,
		"line":         line,
	})

}
