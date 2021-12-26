/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-26 21:00:10
 */
package models

import "database/sql"

type User struct {
	BaseModle
	Uid          uint           `gorm:"primary_key;auto_increment"`
	Nickname     string         `json:"nickname" form:"nickanme"`
	Email        string         `json:"email" form:"email"`
	Password     string         `json:"password" form:"password"`
	Salt         string         `json:"salt" form:"salt"`
	RefreshToken sql.NullString `json:"refresh_token"`
	ExpirTime    sql.NullTime   `json:"expir_time"`
}

type UserRegisterReq struct {
	Nickname        string `json:"nickname" form:"nickname" binding:"required" label:"昵称"`
	Email           string `json:"email" form:"email" binding:"required,email" label:"邮箱"`
	Password        string `json:"password" form:"password" binding:"required" label:"密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password" label:"确认密码"`
}

type UserLoginReq struct {
	Email    string `json:"email" form:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required" label:"密码"`
}

type UserReq struct {
	Username string `json:"username" form:"username" binding:"required" label:"用户名"` //用户名
	Sex      uint   `json:"sex" form:"sex" binding:"required" label:"性别"`            //性别
	Age      uint   `json:"age" form:"age" binding:"required" label:"年龄"`            //年龄
}

type UserLoginRes struct {
	Jtoken  string `json:"jtoken"`  //Jtoken 验证字符串
	Retoken string `json:"retoken"` //retoken 刷新token
}

type UserRefreshTokenReq struct {
	Retoken string `json:"retoken" form:"retoken" binding:"required"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) FillData() {

}
