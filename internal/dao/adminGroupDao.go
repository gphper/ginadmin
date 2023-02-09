package dao

import (
	"sync"

	"github.com/gphper/ginadmin/pkg/mysqlx"

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
		instanceAdminGroup = &AdminGroupDao{DB: mysqlx.GetDB(&mysqlx.BaseModle{ConnName: "default"})}
	})
	return instanceAdminGroup
}
