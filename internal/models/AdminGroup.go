/*
 * @Description:用户组相关model
 * @Author: gphper
 * @Date: 2021-07-11 18:56:22
 */

package models

type AdminGroupSaveReq struct {
	Privs     []string `form:"privs[]" label:"权限" json:"privs" binding:"required"`
	GroupName string   `form:"groupname" label:"用户组名" json:"groupname" binding:"required"`
	GroupId   uint     `form:"groupid"`
}
