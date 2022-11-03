/*
 * @Description:用户组服务
 * @Author: gphper
 * @Date: 2021-07-17 21:03:55
 */
package admin

import (
	"sync"

	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/casbinauth"
)

type adminGroupService struct {
	Dao *dao.AdminGroupDao
}

var (
	instanceAdminGroupService *adminGroupService
	onceAdminGroupService     sync.Once
)

func NewAdminGroupService() *adminGroupService {
	onceAdminGroupService.Do(func() {
		instanceAdminGroupService = &adminGroupService{
			Dao: dao.NewAdminGroupDao(),
		}
	})
	return instanceAdminGroupService
}

//保存角色
func (ser *adminGroupService) SaveGroup(req models.AdminGroupSaveReq) error {
	oldGroup := casbinauth.GetPoliceByGroup(req.GroupName)
	oldLen := len(oldGroup)
	oldSlice := make([]string, oldLen)
	if oldLen > 0 {
		for oldk, oldv := range oldGroup {
			oldSlice[oldk] = oldv[1] + ":" + oldv[2]
		}
	}

	tx := ser.Dao.DB.Begin()

	_, err := casbinauth.UpdatePolices(req.GroupName, oldSlice, req.Privs, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//删除角色
func (ser *adminGroupService) DelGroup(id string) (ok bool, err error) {
	polices := casbinauth.GetPoliceByGroup(id)
	ok, err = casbinauth.DelGroups("p", polices)
	return
}
