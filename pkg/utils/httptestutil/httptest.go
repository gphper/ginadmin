/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-31 20:43:41
 */
package httptestutil

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type OptionValue struct {
	Param   url.Values
	Cookies []*http.Cookie
}

// Get 请求方法
func Get(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, router *gin.Engine, options OptionValue) (body []byte, cookies []*http.Cookie) {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(options.Param.Encode()))

	// 设置cookies
	if len(options.Cookies) > 0 {
		for _, cookie := range options.Cookies {
			req.AddCookie(cookie)
		}
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()

	defer result.Body.Close()

	// 读取响应body
	body, _ = ioutil.ReadAll(result.Body)
	cookies = result.Cookies()
	return
}
