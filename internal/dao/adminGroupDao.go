package dao

import (
	"ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminGroupDao struct {
	DB *gorm.DB
}

var AgDao = adminGroupDao{DB: models.Db}
