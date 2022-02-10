package dao

import (
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminUserDao struct {
	DB *gorm.DB
	Tx *gorm.DB
}

var AuDao = adminUserDao{DB: models.Db, Tx: models.Db.Begin()}
