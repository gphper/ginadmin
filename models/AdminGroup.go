package models

import "time"

type AdminGroup struct {
	GroupId uint `gorm:"primary_key;auto_increment"`
	GroupName string `gorm:"size:20;comment:'用户组名称'"`
	Privs string `gorm:"type:text;comment:'权限'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AdminGroup) TableName() string {
	return "admin_groups"
}