/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-08 20:12:04
 */
package models

func GetModels() []interface{} {
	return []interface{}{
		&AdminUsers{}, &Article{}, &UploadType{}, &User{},
	}
}
