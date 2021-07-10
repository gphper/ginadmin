package models

import "time"

type AdminGroup struct {
	BaseModle
	GroupId   uint      `gorm:"primary_key;auto_increment"`
	GroupName string    `gorm:"size:20;comment:'用户组名称'"`
	Privs     string    `gorm:"type:text;comment:'权限'"`
	CreatedAt time.Time `gorm:"size:0"`
	UpdatedAt time.Time `gorm:"size:0"`
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

func (ag *AdminGroup) TableName() string {
	return "admin_groups"
}

func (ag *AdminGroup) FillData() {
	//填充管理用户组
	adminGroup := AdminGroup{
		GroupId:   1,
		GroupName: "管理员组",
		Privs:     "{\"all\":{}}",
	}
	Db.Save(&adminGroup)
}
