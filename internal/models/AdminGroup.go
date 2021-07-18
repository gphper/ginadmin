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

type AdminGroupSaveReq struct {
	Privs     []string `form:"privs[]" label:"权限" json:"privs" binding:"required"`
	GroupName string   `form:"groupname" label:"用户组名" json:"groupname" binding:"required"`
	GroupId   uint     `form:"groupid"`
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
