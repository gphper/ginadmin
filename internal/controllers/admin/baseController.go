/*
 * @Description:基础控制器
 * @Author: gphper
 * @Date: 2021-05-18 23:21:06
 */

package admin

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gphper/ginadmin/internal/errorx"
	"github.com/gphper/ginadmin/pkg/loggers"
	gvalidator "github.com/gphper/ginadmin/pkg/validator"
	perrors "github.com/pkg/errors"

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
		m := fmt.Sprintf("%+v", err)
		m = strings.ReplaceAll(m, "\n", "<br/>")
		m = strings.ReplaceAll(m, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;")
		c.HTML(http.StatusOK, "home/debug.html", gin.H{
			"title": "DEBUG",
			"msg":   template.HTML(m),
			"error": err.Error(),
		})
	} else {
		//收集信息存入到日志
		ctx, _ := c.Get("ctx")
		var code int
		var msg string

		sourceErr := perrors.Cause(err)
		customErr, ok := sourceErr.(*errorx.CustomError)
		if !ok {
			code = errorx.HTTP_UNKNOW_ERR
			msg = err.Error()
		} else {
			code = customErr.ErrCode
			msg = customErr.ErrMsg

			if customErr.Err != nil {
				//只记录带有wrap的error
				loggers.LogError(ctx.(context.Context), "admin-custom-error", "err msg", map[string]string{
					"err msg": err.Error(),
					"stack":   fmt.Sprintf("%+v", err),
				})
			}
		}

		c.HTML(http.StatusOK, "home/error.html", gin.H{
			"title": "出错了~",
			"code":  code,
			"msg":   msg,
		})
	}
}

func (Base BaseController) Html(c *gin.Context, code int, name string, data gin.H) {

	if data == nil {
		data = gin.H{}
	}

	uid, _ := c.Get("uid")
	username, _ := c.Get("username")
	groupname, _ := c.Get("groupname")
	data["username"] = username
	data["uid"] = uid
	data["groupname"] = groupname

	c.HTML(code, name, data)
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
