package dao

import (
	"sync"

	"github.com/gphper/ginadmin/internal/models"

	"gorm.io/gorm"
)

type AdminGroupDao struct {
	DB *gorm.DB
}

var (
	instanceAdminGroup *AdminGroupDao
	onceAdminGroup     sync.Once
)

func NewAdminGroupDao() *AdminGroupDao {
	onceAdminGroup.Do(func() {
		instanceAdminGroup = &AdminGroupDao{DB: models.GetDB(&models.BaseModle{ConnName: "default"})}
	})
	return instanceAdminGroup
}
