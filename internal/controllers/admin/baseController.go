/*
 * @Description:基础控制器
 * @Author: gphper
 * @Date: 2021-05-18 23:21:06
 */

package admin

import (
	"errors"
	"net/http"
	"runtime"
	"strconv"

	gvalidator "github.com/gphper/ginadmin/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseController struct {
}

func (Base BaseController) Success(c *gin.Context, url string, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"msg":         message,
		"url":         url,
		"iframe_jump": false,
	})
}

func (Base BaseController) Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": false,
		"msg":    message,
	})
}

func (Base BaseController) ErrorHtml(c *gin.Context, err error) {

	if gin.Mode() == "debug" {
		_, file, line, _ := runtime.Caller(1)
		c.HTML(http.StatusOK, "home/debug.html", gin.H{
			"title": "DEBUG",
			"msg":   file + ":" + strconv.Itoa(line),
			"error": err.Error(),
		})
	} else {
		c.HTML(http.StatusOK, "home/error.html", gin.H{
			"title": "出错了~",
		})
	}
}

func (Base BaseController) FormBind(c *gin.Context, obj interface{}) error {

	trans, err := gvalidator.InitTrans("zh")

	if err != nil {
		return err
	}

	if err := c.ShouldBind(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok && errs != nil {
			return errs
		}

		for _, v := range errs.Translate(trans) {
			return errors.New(v)
		}

	}
	return nil
}

func (Base BaseController) UriBind(c *gin.Context, obj interface{}) error {

	trans, err := gvalidator.InitTrans("zh")

	if err != nil {
		return err
	}

	if err := c.ShouldBindUri(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok && errs != nil {
			return errs
		}

		for _, v := range errs.Translate(trans) {
			return errors.New(v)
		}

	}
	return nil
}
