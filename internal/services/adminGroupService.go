package services

import (
	"encoding/json"
	"ginadmin/internal/dao"
	"ginadmin/internal/models"
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
	var privsJsonStr string
	privMap := make(map[string]struct{})
	//将数组转为map便于提高后面的判断效率
	for _, v := range req.Privs {
		privMap[v] = struct{}{}
	}

	privsJson, err := json.Marshal(privMap)
	if err == nil {
		privsJsonStr = string(privsJson)
	} else {
		privsJsonStr = `[]`
	}

	adminGroup := models.AdminGroup{
		GroupId:   req.GroupId,
		GroupName: req.GroupName,
		Privs:     privsJsonStr,
	}
	return dao.AgDao.DB.Save(&adminGroup).Error
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
func (ser *adminGroupService) DelGroup(id string) (err error) {
	err = dao.AgDao.DB.Where("group_id = ?", id).Delete(models.AdminGroup{}).Error
	return
}
