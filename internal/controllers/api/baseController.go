/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-15 21:51:07
 */
package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gphper/ginadmin/internal/errorx"
	"github.com/gphper/ginadmin/pkg/loggers"
	gvalidator "github.com/gphper/ginadmin/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	perrors "github.com/pkg/errors"
)

type BaseController struct {
}

type SuccessResponse struct {
	DefaultResponse
	Data interface{} `json:"data" swaggertype:"object"` //接口返回的业务数据
}

type DefaultResponse struct {
	Code int    `json:"code"` //code 为1表示正常 0表示业务请求错误
	Msg  string `json:"msg"`  //错误提示信息
}

func (Base BaseController) Success(c *gin.Context, obj interface{}) {
	var res SuccessResponse
	res.Code = 200
	res.Msg = "success"
	res.Data = obj

	c.JSON(http.StatusOK, res)
}

func (Base BaseController) Error(c *gin.Context, err error) {

	var res DefaultResponse

	sourceErr := perrors.Cause(err)
	customErr, ok := sourceErr.(*errorx.CustomError)
	if ok {
		res.Code = customErr.ErrCode
		res.Msg = customErr.ErrMsg
		// 保存日志
		if customErr.Err != nil {
			ctx, _ := c.Get("ctx")
			loggers.LogError(ctx.(context.Context), "api-custom-error", "err msg", map[string]string{
				"err msg": err.Error(),
				"stack":   fmt.Sprintf("%+v", err),
			})
		}
	} else {
		res.Code = errorx.HTTP_UNKNOW_ERR
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res)
}

func (Base BaseController) FormBind(c *gin.Context, obj interface{}) error {

	trans, err := gvalidator.InitTrans("en")

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
		return err

	}
	return nil
}
