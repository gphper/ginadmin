/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-22 09:52:35
 */
package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

var (
	instanceUser *UserDao
	onceUserDao  sync.Once
)

func NewUserDao() *UserDao {
	onceUserDao.Do(func() {
		instanceUser = &UserDao{DB: models.Db}
	})
	return instanceUser
}

func (dao *UserDao) GetUser(conditions map[string]interface{}) (user models.User, err error) {
	err = dao.DB.Where(conditions).First(&user).Error
	return
}

func (dao *UserDao) UpdateColumns(conditions, field map[string]interface{}, tx *gorm.DB) error {

	if tx != nil {
		return tx.Model(&models.User{}).Where(conditions).UpdateColumns(field).Error
	}

	return dao.DB.Model(&models.User{}).Where(conditions).UpdateColumns(field).Error
}
