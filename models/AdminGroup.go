package models

import "time"

type AdminGroup struct {
	GroupId   uint   `gorm:"primary_key;auto_increment"`
	GroupName string `gorm:"size:20;comment:'用户组名称'"`
	Privs     string `gorm:"type:text;comment:'权限'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AdminGroup) TableName() string {
	return "admin_groups"
}

func GetAllAdminGroup() ([]AdminGroup, error) {
	var adminGroups []AdminGroup
	error := Db.Where("group_id != ?", 1).Find(&adminGroups).Error
	return adminGroups, error
}

func SaveAdminGroup(grouid uint, groupname string, privsJsonStr string) error {
	adminGroup := AdminGroup{
		GroupId:   grouid,
		GroupName: groupname,
		Privs:     privsJsonStr,
	}
	return Db.Save(&adminGroup).Error
}

func FindAdminGroupById(id string) (AdminGroup, error) {
	var adminGroup AdminGroup
	err := Db.Where("group_id = ?", id).First(&adminGroup).Error
	return adminGroup, err
}

func DelAdminGroupById(id string) error {
	return Db.Where("group_id = ?", id).Delete(AdminGroup{}).Error
}
