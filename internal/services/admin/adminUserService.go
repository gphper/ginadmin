/*
 * @Description:用户服务
 * @Author: gphper
 * @Date: 2021-07-18 13:59:07
 */
package admin

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/casbinauth"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"

	"gorm.io/gorm"
)

type adminUserService struct {
	Dao *dao.AdminUserDao
}

var (
	instanceAdminUserService *adminUserService
	onceAdminUserService     sync.Once
)

func NewAdminUserService() *adminUserService {
	onceAdminUserService.Do(func() {
		instanceAdminUserService = &adminUserService{
			Dao: dao.NewAdminUserDao(),
		}
	})
	return instanceAdminUserService
}

//获取管理员
func (ser *adminUserService) GetAdminUsers(req models.AdminUserIndexReq) (db *gorm.DB) {
	return ser.Dao.GetAdminUsers(req.Nickname, req.CreatedAt)
}

//添加或保存管理员信息
func (ser *adminUserService) SaveAdminUser(req models.AdminUserSaveReq) (err error) {

	var (
		adminUser models.AdminUsers
		ok        bool
	)
	groupnameStr, _ := json.Marshal(req.GroupName)

	var rules = make([][]string, 0)
	for _, v := range req.GroupName {
		rules = append(rules, []string{req.Username, v})
	}

	tx := ser.Dao.DB.Begin()

	if req.Uid > 0 {

		var groupOldName []string
		adminUser, err = ser.Dao.GetAdminUser(map[string]interface{}{"uid": req.Uid})
		if err != nil {
			return
		}

		json.Unmarshal([]byte(adminUser.GroupName), &groupOldName)
		fields := map[string]interface{}{
			"group_name": string(groupnameStr),
			"username":   req.Username,
			"nickname":   req.Nickname,
			"phone":      req.Phone,
		}
		if req.Password != "" {
			salt := gstrings.RandString(6)
			fields["salt"] = salt
			fields["password"] = gstrings.Encryption(req.Password, salt)
		}

		err = ser.Dao.UpdateColumns(map[string]interface{}{"uid": req.Uid}, fields, tx)
		if err != nil {
			tx.Rollback()
			return
		}

		_, err = casbinauth.UpdateGroups(req.Username, groupOldName, req.GroupName, tx)
		if err != nil {
			tx.Rollback()
			return
		}

	} else {
		salt := gstrings.RandString(6)
		passwordSalt := gstrings.Encryption(req.Password, salt)
		adminUser := models.AdminUsers{
			GroupName: string(groupnameStr),
			Nickname:  req.Nickname,
			Username:  req.Username,
			Password:  passwordSalt,
			Phone:     req.Phone,
			Salt:      salt,
		}
		err = tx.Save(&adminUser).Error
		if err != nil {
			tx.Rollback()
			return
		}
		//将权限添加到casbin中
		ok, err = casbinauth.AddGroups("g", rules, tx)
		if err != nil || !ok {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return
}

//获取单个管理员用户信息
func (ser *adminUserService) GetAdminUser(conditions map[string]interface{}) (adminUser models.AdminUsers, err error) {
	adminUser, err = ser.Dao.GetAdminUser(conditions)
	return
}

//删除管理员
func (ser *adminUserService) DelAdminUser(id string) (err error) {
	return ser.Dao.Del(map[string]interface{}{"uid": id})
}

//修改密码
func (ser *adminUserService) EditPass(req models.AdminUserEditPassReq) (err error) {

	var adminUser models.AdminUsers

	if req.NewPassword != req.SubPassword {
		err = errors.New("请再次确认新密码是否正确")
		return
	}

	adminUser, err = ser.GetAdminUser(map[string]interface{}{"uid": req.Uid})
	if err != nil {
		return
	}

	oldPass := gstrings.Encryption(req.OldPassword, adminUser.Salt)
	if oldPass != adminUser.Password {
		err = errors.New("原密码错误")
		return
	}

	newPass := gstrings.Encryption(req.NewPassword, adminUser.Salt)
	err = ser.Dao.UpdateColumn(adminUser.Uid, "password", newPass)

	return
}

//根究用户保存自定义皮肤
func (ser *adminUserService) EditSkin(req models.AdminUserSkinReq) (err error) {

	var skinMap = map[string]string{
		"data-logobg":    "logo",
		"data-sidebarbg": "side",
		"data-headerbg":  "header",
	}

	err = ser.Dao.UpdateColumn(uint(req.Uid), skinMap[req.Type], req.Color)

	return
}
