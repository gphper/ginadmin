package dao

import (
	"github/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminGroupDao struct {
	DB *gorm.DB
}

var AgDao = adminGroupDao{DB: models.Db}
