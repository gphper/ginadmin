/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-15 20:14:09
 */
package dao

import (
	"context"
	"strings"
	"sync"

	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/mysqlx"

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
		instanceAdminUser = &AdminUserDao{DB: mysqlx.GetDB(&models.AdminUsers{})}
	})
	return instanceAdminUser
}

func (dao *AdminUserDao) GetAdminUser(conditions map[string]interface{}) (adminUser models.AdminUsers, err error) {
	err = dao.DB.Where(conditions).First(&adminUser).Error
	return
}

func (dao *AdminUserDao) GetAdminUsers(context context.Context, nickname string, created_time string) (db *gorm.DB) {

	db = dao.DB.WithContext(context).Table("admin_users")

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

func (dao *AdminUserDao) UpdateColumns(conditions, field map[string]interface{}, tx *gorm.DB) error {

	if tx != nil {
		return tx.Model(&models.AdminUsers{}).Where(conditions).UpdateColumns(field).Error
	}

	return dao.DB.Model(&models.AdminUsers{}).Where(conditions).UpdateColumns(field).Error
}

func (dao *AdminUserDao) Del(conditions map[string]interface{}) error {
	return dao.DB.Delete(&models.AdminUsers{}, conditions).Error
}
