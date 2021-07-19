package models

import (
	"ginadmin/pkg/comment"
	"time"

	"gorm.io/gorm"
)

type AdminUsers struct {
	BaseModle
	Uid       uint   `gorm:"primary_key;auto_increment"`
	GroupId   uint   `gorm:"size:20;comment:'用户组id'"`
	Username  string `gorm:"size:100;comment:'用户名'"`
	Nickname  string `gorm:"size:100;comment:'姓名'"`
	Password  string `gorm:"size:200;comment:'密码'"`
	Phone     string `gorm:"size:20;comment:'手机号'"`
	LastLogin string `gorm:"size:30;comment:'最后登录ip地址'"`
	Salt      string `gorm:"size:32;comment:'密码盐'"`
	ApiToken  string `gorm:"size:32;comment:'用户登录凭证'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminUserInfo struct {
	AdminUsers
	GroupName string
}

type AdminUserIndexReq struct {
	Nickname  string `form:"nickname"`
	CreatedAt string `form:"created_at"`
}

type AdminUserSaveReq struct {
	Username string `form:"username" label:"用户名" binding:"required"`
	Password string `form:"password"`
	Nickname string `form:"nickname" label:"姓名" binding:"required"`
	Phone    string `form:"phone"`
	GroupId  uint   `form:"groupid" label:"用户组" binding:"required"`
	Uid      uint   `form:"uid"`
}

func GetAllAdminUserJoinGroup() *gorm.DB {
	return Db.Table("admin_users").Joins("join admin_groups on admin_groups.group_id = admin_users.group_id").Select("admin_users.*,admin_groups.group_name").Where("uid != ?", 1)
}

func GetAllAdminUserJoinGroupLikeNickname(db *gorm.DB, nickname string) *gorm.DB {
	return db.Where("nickname like ?", "%"+nickname+"%")
}

func GetAllAdminUserJoinGroupTimeRange(db *gorm.DB, start string, end string) *gorm.DB {
	return db.Where("admin_users.created_at > ? ", start).Where("admin_users.created_at < ?", end)
}

func AddAdminUser(groupid int, username string, nickname string, phone string, password string) error {
	salt := comment.RandString(6)
	passwordSalt := comment.Encryption(password, salt)
	adminUser := AdminUsers{
		GroupId:   uint(groupid),
		Username:  username,
		Nickname:  nickname,
		Password:  passwordSalt,
		Phone:     phone,
		LastLogin: "",
		Salt:      salt,
		ApiToken:  "",
	}
	return Db.Save(&adminUser).Error
}

func SaveAdminUser(uid int, groupid int, nickname string, phone string, password string) error {
	adminUser := AdminUsers{
		Uid:       uint(uid),
		GroupId:   uint(groupid),
		Nickname:  nickname,
		Password:  "",
		Phone:     phone,
		LastLogin: "",
		Salt:      "",
		ApiToken:  "",
	}
	if password != "" {
		salt := comment.RandString(6)
		adminUser.Salt = salt
		passwordSalt := comment.Encryption(password, salt)
		adminUser.Password = passwordSalt
	}
	return Db.Model(&adminUser).Updates(adminUser).Error
}

func GetAdminUserById(id string) (AdminUsers, error) {
	var adminUser AdminUsers
	err := Db.Table("admin_users").Where("uid = ?", id).First(&adminUser).Error
	return adminUser, err
}

func AlterAdminUserPass(id string, password string) error {
	return Db.Table("admin_users").Where("uid = ?", id).Updates(map[string]interface{}{
		"password": password,
	}).Error
}

func (au *AdminUsers) TableName() string {
	return "admin_users"
}

func (au *AdminUsers) FillData() {
	//初始化管理员
	salt := comment.RandString(6)
	passwordSalt := comment.Encryption("111111", salt)
	adminUser := AdminUsers{
		Uid:       1,
		GroupId:   1,
		Username:  "admin",
		Nickname:  "管理员",
		Password:  passwordSalt,
		Phone:     "",
		LastLogin: "",
		Salt:      salt,
		ApiToken:  "",
	}
	Db.Save(&adminUser)
}