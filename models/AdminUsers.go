package models

import "time"

type AdminUsers struct {
	Uid uint `gorm:"primary_key;auto_increment"`
	GroupId uint `gorm:"size:20;comment:'用户组id'"`
	Username string `gorm:"size:100;comment:'用户名'"`
	Nickname string `gorm:"size:100;comment:'姓名'"`
	Password string `gorm:"size:200;comment:'密码'"`
	Phone string `gorm:"size:20;comment:'手机号'"`
	LastLogin string `gorm:"size:30;comment:'最后登录ip地址'"`
	Salt string `gorm:"size:32;comment:'密码盐'"`
	ApiToken string `gorm:"size:32;comment:'用户登录凭证'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}