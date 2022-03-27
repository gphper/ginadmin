/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-22 09:55:12
 */
package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/jwt"
	"github.com/gphper/ginadmin/pkg/utils/strings"
)

type userService struct{}

var UserService = userService{}

/**
* 用户注册
**/
func (ser *userService) Register(req models.UserRegisterReq) error {
	var (
		user models.User
		err  error
	)

	err = dao.UserDao.DB.Find(&user, "email = ?", req.Email).Error
	if err != nil {
		return err
	}

	if user.Uid != 0 {
		return errors.New("该邮箱已存在")
	}

	user.Nickname = req.Nickname
	user.Email = req.Email
	salt := strings.RandString(6)
	passwordSalt := strings.Encryption(req.Password, salt)
	user.Password = passwordSalt
	user.Salt = salt

	return dao.UserDao.DB.Create(&user).Error
}

/**
* 验证用户登录
 */
func (ser *userService) Login(req models.UserLoginReq) (jtoken string, retoken string, err error) {
	var (
		user    models.User
		payload jwt.Payload
	)

	err = dao.UserDao.DB.Find(&user, "email = ?", req.Email).Error
	if err != nil {
		return
	}

	if user.Uid == 0 {
		return jtoken, retoken, errors.New("账号密码错误")
	}

	//校验密码
	passwordSalt := strings.Encryption(req.Password, user.Salt)
	if passwordSalt != user.Password {
		return jtoken, retoken, errors.New("账号密码错误")
	}

	//生成jtoken
	payload.Name = user.Nickname
	payload.Uid = user.Uid
	payload.Exp = time.Now().Local().Add(5 * time.Minute)
	jtoken, err = jwt.Generate("HS256", payload)
	if err != nil {
		return
	}

	//生成refresh_token
	retoken = strings.Encryption(passwordSalt, strconv.FormatInt(time.Now().UnixNano(), 10))
	user.RefreshToken.String = retoken
	user.RefreshToken.Valid = true
	user.ExpirTime.Time = time.Now().Add(7 * time.Hour)
	dao.UserDao.DB.Save(&user)

	return
}

/**
* 使用refresh token 更换jtoken
 */
func (ser *userService) RefreshToken(req models.UserRefreshTokenReq) (jtoken string, err error) {
	var (
		user    models.User
		payload jwt.Payload
	)

	err = dao.UserDao.DB.Find(&user, "refresh_token = ?", req.Retoken).Error
	if err != nil {
		return
	}

	if user.Uid == 0 {
		return jtoken, errors.New("refresh token 错误")
	}

	//校验过期时间
	if time.Until(user.ExpirTime.Time).Hours() < 0 {
		return jtoken, errors.New("refresh token 已过期请重新登录")
	}

	//生成jtoken
	payload.Name = user.Nickname
	payload.Uid = user.Uid
	payload.Exp = time.Now().Local().Add(5 * time.Minute)
	jtoken, err = jwt.Generate("HS256", payload)
	if err != nil {
		return jtoken, err
	}

	return
}
