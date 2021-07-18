package services

import (
	"errors"
	"ginadmin/internal/dao"
	"ginadmin/internal/models"
	"ginadmin/pkg/comment"
	"strings"

	"gorm.io/gorm"
)

type adminUserService struct{}

var AuService = adminUserService{}

//获取管理员
func (ser *adminUserService) GetAdminUsers(req models.AdminUserIndexReq) (db *gorm.DB) {

	db = dao.AuDao.DB.Table("admin_users").Joins("join admin_groups on admin_groups.group_id = admin_users.group_id").Select("admin_users.*,admin_groups.group_name").Where("uid != ?", 1)

	if req.Nickname != "" {
		db = db.Where("nickname like ?", "%"+req.Nickname+"%")
	}

	if req.CreatedAt != "" {
		period := strings.Split(req.CreatedAt, " ~ ")
		start := period[0] + " 00:00:00"
		end := period[1] + " 23:59:59"
		db = db.Where("admin_users.created_at > ? ", start).Where("admin_users.created_at < ?", end)
	}

	return
}

//添加或保存管理员信息
func (ser *adminUserService) SaveAdminUser(req models.AdminUserSaveReq) error {
	if req.Uid == 0 {
		if len(req.Password) == 0 {
			return errors.New("请填写密码")
		}
		var count int64
		//验证用户唯一性
		dao.AuDao.DB.Table("admin_users").Where("username", req.Username).Count(&count)
		if count != 0 {
			return errors.New("当前用户名已存在")
		}
		salt := comment.RandString(6)
		passwordSalt := comment.Encryption(req.Password, salt)
		adminUser := models.AdminUsers{
			GroupId:   req.GroupId,
			Username:  req.Username,
			Nickname:  req.Nickname,
			Password:  passwordSalt,
			Phone:     req.Phone,
			LastLogin: "",
			Salt:      salt,
			ApiToken:  "",
		}
		return dao.AuDao.DB.Save(&adminUser).Error
	} else {
		adminUser := models.AdminUsers{
			Uid:       req.Uid,
			GroupId:   req.GroupId,
			Nickname:  req.Nickname,
			Password:  "",
			Phone:     req.Phone,
			LastLogin: "",
			Salt:      "",
			ApiToken:  "",
		}
		if len(req.Password) != 0 {
			salt := comment.RandString(6)
			adminUser.Salt = salt
			passwordSalt := comment.Encryption(req.Password, salt)
			adminUser.Password = passwordSalt
		}
		return dao.AuDao.DB.Model(&adminUser).Updates(adminUser).Error
	}

}

//获取单个管理员用户信息
func (ser *adminUserService) GetAdminUser(id string) (adminUser models.AdminUsers, err error) {
	err = dao.AuDao.DB.Where("uid = ?", id).First(&adminUser).Error
	return
}

//删除管理员
func (ser *adminUserService) DelAdminUser(id string) (err error) {
	err = dao.AuDao.DB.Where("uid = ?", id).Delete(models.AdminUsers{}).Error
	return
}
