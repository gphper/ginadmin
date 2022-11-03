/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-15 20:14:09
 */
package dao

import (
	"strings"
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type AdminUserDao struct {
	DB *gorm.DB
}

var (
	instanceAdminUser *AdminUserDao
	onceAdminUserDao  sync.Once
)

func NewAdminUserDao() *AdminUserDao {
	onceAdminUserDao.Do(func() {
		instanceAdminUser = &AdminUserDao{DB: models.Db}
	})
	return instanceAdminUser
}

func (dao *AdminUserDao) GetAdminUser(conditions map[string]interface{}) (adminUser models.AdminUsers, err error) {
	err = dao.DB.Where(conditions).First(&adminUser).Error
	return
}

func (dao *AdminUserDao) GetAdminUsers(uid int, nickname string, created_time string) (db *gorm.DB) {

	db = dao.DB.Table("admin_users").Where("uid != ?", uid)

	if nickname != "" {
		db = db.Where("nickname like ?", "%"+nickname+"%")
	}

	if created_time != "" {
		period := strings.Split(created_time, " ~ ")
		start := period[0] + " 00:00:00"
		end := period[1] + " 23:59:59"
		db = db.Where("admin_users.created_at > ? ", start).Where("admin_users.created_at < ?", end)
	}

	return
}

func (dao *AdminUserDao) UpdateColumn(uid uint, key, value string) error {
	return dao.DB.Model(&models.AdminUsers{}).Where("uid = ?", uid).UpdateColumn(key, value).Error
}
