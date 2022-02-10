package dao

import (
	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminGroupDao struct {
	DB *gorm.DB
}

var AgDao = adminGroupDao{DB: models.Db}
