/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-20 20:59:47
 */
package models

type UploadType struct {
	BaseModle
	Id          uint   `gorm:"primary_key;auto_increment"`
	TypeName    string `gorm:"size:100;comment:'类型名称'"`
	StoragePath string `gorm:"size:100;comment:'存储路径'"`
	Description string `gorm:"size:100;comment:'描述'"`
	AllowType   string `gorm:"size:100;comment:'允许上传的资源类型'"`
	AllowSize   uint   `gorm:"comment:'允许上传的资源大小，单位KB'"`
	AllowNum    uint   `gorm:"comment:'允许上传的资源个数'"`
}

type UploadHtmlReq struct {
	TypeName string `uri:"type_name"`
	Id       string `uri:"id"`
	Type     uint   `uri:"type"`
	NowNum   uint   `uri:"now_num"`
}

func (u *UploadType) TableName() string {
	return "upload_type"
}

func (u *UploadType) FillData() {
	ut := UploadType{
		Id:          1,
		TypeName:    "UP_COMMON",
		StoragePath: "common",
		Description: "基础资源存储",
		AllowType:   "jpg|png|jpeg|gif",
		AllowSize:   5242880,
		AllowNum:    1,
	}
	Db.Save(&ut)
}
