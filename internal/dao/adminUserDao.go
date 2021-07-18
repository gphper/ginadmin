package dao

import (
	"ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminUserDao struct {
	DB *gorm.DB
}

var AuDao = adminUserDao{DB: models.Db}
