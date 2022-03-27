/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-15 21:51:07
 */
package api

import (
	"errors"
	"net/http"

	gvalidator "github.com/gphper/ginadmin/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseController struct {
}

type SuccessResponse struct {
	DefaultResponse
	Data interface{} `json:"data" swaggertype:"object"` //接口返回的业务数据
}

type DefaultResponse struct {
	Code uint   `json:"code"` //code 为1表示正常 0表示业务请求错误
	Msg  string `json:"msg"`  //错误提示信息
}

func (Base BaseController) Success(c *gin.Context, obj interface{}) {
	var res SuccessResponse
	res.Code = 1
	res.Msg = "success"
	res.Data = obj

	c.JSON(http.StatusOK, res)
}

func (Base BaseController) Error(c *gin.Context, err error) {
	var res DefaultResponse
	res.Code = 0
	res.Msg = err.Error()

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
