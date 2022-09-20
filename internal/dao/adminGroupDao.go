package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type adminGroupDao struct {
	DB *gorm.DB
}

var insAgd *adminGroupDao
var onceAgd sync.Once

func NewAdminGroupDao() *adminGroupDao {
	onceAgd.Do(func() {
		insAgd = &adminGroupDao{DB: models.Db}
	})
	return insAgd
}
