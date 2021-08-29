/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-28 21:37:22
 */
package store

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type sessionStore struct {
	session sessions.Session
	expir   int
}

func NewSessionStore(c *gin.Context, time int) base64Captcha.Store {
	s := new(sessionStore)
	s.session = sessions.Default(c)
	s.expir = time
	return s
}

func (s *sessionStore) Set(id string, value string) error {

	s.session.Set("captcha", value)
	s.session.Set("captcha_time", time.Now().Unix())
	s.session.Save()
	return nil
}

func (s *sessionStore) Verify(id, answer string, clear bool) bool {
	val := s.session.Get("captcha").(string)
	time_before := s.session.Get("captcha_time")

	v, _ := time_before.(int64)

	if time.Now().Unix()-v > int64(s.expir) {
		return false
	}

	return answer == val
}

func (s *sessionStore) Get(id string, clear bool) (value string) {
	return s.session.Get("captcha").(string)
}
