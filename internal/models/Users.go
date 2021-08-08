/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-26 21:00:10
 */
package models

type UserReq struct {
	Username string `json:"username" form:"username" binding:"required" label:"用户名"` //用户名
	Sex      uint   `json:"sex" form:"sex" binding:"required" label:"性别"`            //性别
	Age      uint   `json:"age" form:"age" binding:"required" label:"年龄"`            //年龄
}
