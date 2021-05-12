package setting

import (
	"bufio"
	"fmt"
	"ginadmin/comment"
	"ginadmin/comment/loggers"
	"ginadmin/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type AdminSystemController struct {
	controllers.BaseController
}

/**
日志目录页面
*/
func (con *AdminSystemController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {

		path, err := comment.RootPath()
		if err != nil {
			fmt.Printf("get root path err:%v", err)
		}
		path = path + comment.GetLine() + "logs"

		files, err := ioutil.ReadDir(path)
		if err != nil {
			loggers.AdminLogger.Error("读取目录失败", zap.Error(err))
		}
		c.HTML(http.StatusOK, "setting/systemlog.html", gin.H{
			"log_path": path,
			"files":    files,
			"line":     comment.GetLine(),
		})
	}
}

/**
获取目录
*/
func (con *AdminSystemController) GetDir() gin.HandlerFunc {
	return func(c *gin.Context) {
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
				Path: path + comment.GetLine() + v.Name(),
				Type: fileType,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": fileSlice,
		})
	}
}

/**
获取日志详情
*/
func (con *AdminSystemController) View() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
