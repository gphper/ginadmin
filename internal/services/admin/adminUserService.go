/*
 * @Description:用户服务
 * @Author: gphper
 * @Date: 2021-07-18 13:59:07
 */
package admin

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/gphper/ginadmin/internal/dao"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/pkg/casbinauth"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"

	"gorm.io/gorm"
)

type adminUserService struct{}

var AuService = adminUserService{}

//获取管理员
func (ser *adminUserService) GetAdminUsers(req models.AdminUserIndexReq) (db *gorm.DB) {

	db = dao.AuDao.DB.Table("admin_users").Where("uid != ?", 1)

	if req.Nickname != "" {
		db = db.Where("nickname like ?", "%"+req.Nickname+"%")
	}

	if req.CreatedAt != "" {
		period := strings.Split(req.CreatedAt, " ~ ")
		start := period[0] + " 00:00:00"
		end := period[1] + " 23:59:59"
		db = db.Where("admin_users.created_at > ? ", start).Where("admin_users.created_at < ?", end)
	}

	return
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

	tx := dao.AuDao.DB.Begin()

	if req.Uid > 0 {
		err = tx.Table("admin_users").Where("uid = ?", req.Uid).First(&adminUser).Error
		if err != nil {
			return
		}

		var groupOldName []string
		json.Unmarshal([]byte(adminUser.GroupName), &groupOldName)

		adminUser := models.AdminUsers{
			Uid:       req.Uid,
			GroupName: string(groupnameStr),
			Username:  req.Username,
			Nickname:  req.Nickname,
			Password:  "",
			Phone:     req.Phone,
			LastLogin: "",
			Salt:      "",
			ApiToken:  "",
		}
		if req.Password != "" {
			salt := gstrings.RandString(6)
			adminUser.Salt = salt
			passwordSalt := gstrings.Encryption(req.Password, salt)
			adminUser.Password = passwordSalt
		}
		err = tx.Model(&adminUser).Updates(adminUser).Error

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
func (ser *adminUserService) GetAdminUser(id string) (adminUser models.AdminUsers, err error) {
	err = dao.AuDao.DB.Where("uid = ?", id).First(&adminUser).Error
	return
}

//删除管理员
func (ser *adminUserService) DelAdminUser(id string) (err error) {
	err = dao.AuDao.DB.Where("uid = ?", id).Delete(models.AdminUsers{}).Error
	return
}

//修改密码
func (ser *adminUserService) EditPass(req models.AdminUserEditPassReq) (err error) {

	var adminUser models.AdminUsers

	if req.NewPassword != req.SubPassword {
		err = errors.New("请再次确认新密码是否正确")
		return
	}

	adminUser, err = ser.GetAdminUser(strconv.Itoa(req.Uid))
	if err != nil {
		return
	}

	oldPass := gstrings.Encryption(req.OldPassword, adminUser.Salt)
	if oldPass != adminUser.Password {
		err = errors.New("原密码错误")
		return
	}

	newPass := gstrings.Encryption(req.NewPassword, adminUser.Salt)

	err = dao.AuDao.DB.Model(&adminUser).UpdateColumn("password", newPass).Error

	return
}

//根究用户保存自定义皮肤
func (ser *adminUserService) EditSkin(req models.AdminUserSkinReq) (err error) {
	var adminUser models.AdminUsers
	var skinMap = map[string]string{
		"data-logobg":    "logo",
		"data-sidebarbg": "side",
		"data-headerbg":  "header",
	}

	adminUser, err = ser.GetAdminUser(strconv.Itoa(req.Uid))
	if err != nil {
		return
	}

	err = dao.AuDao.DB.Model(&adminUser).UpdateColumn(skinMap[req.Type], req.Color).Error

	return
}
