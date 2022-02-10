/*
 * @Description:用户组服务
 * @Author: gphper
 * @Date: 2021-07-17 21:03:55
 */
package admin

import (
	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/casbinauth"
)

type adminGroupService struct{}

var AgService = adminGroupService{}

//获取角色列表
func (ser *adminGroupService) GetList() (adminGroups []models.AdminGroup, err error) {
	err = dao.AgDao.DB.Where("group_id != ?", 1).Find(&adminGroups).Error
	return
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

	tx := models.Db.Begin()

	_, err := casbinauth.UpdatePolices(req.GroupName, oldSlice, req.Privs, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

//获取角色信息
func (ser *adminGroupService) GetGroup(id string) (adminGroup models.AdminGroup, err error) {
	err = dao.AgDao.DB.Where("group_id = ?", id).First(&adminGroup).Error
	return
}

//获取所有角色
func (ser *adminGroupService) GetAllGroup() (adminGroups []models.AdminGroup, err error) {
	err = dao.AgDao.DB.Where("group_id != ?", 1).Find(&adminGroups).Error
	return
}

//删除角色
func (ser *adminGroupService) DelGroup(id string) (ok bool, err error) {
	polices := casbinauth.GetPoliceByGroup(id)
	ok, err = casbinauth.DelGroups("p", polices)
	return
}
