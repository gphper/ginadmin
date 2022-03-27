/*
 * @Description:casbin方法二次封装
 * @Author: gphper
 * @Date: 2021-07-20 20:52:20
 */
package casbinauth

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gphper/ginadmin/internal/models"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var syncPool sync.Pool

func init() {
	syncPool = sync.Pool{
		New: func() interface{} {
			a, err := gormadapter.NewAdapterByDB(models.Db)
			if err != nil {
				fmt.Println(err)
			}

			text :=
				`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
			m, _ := model.NewModelFromString(text)

			e, _ := casbin.NewEnforcer(m, a)
			return e
		},
	}
}

//获取enforce对象，开始事务时使用
func newEnforceObj(tx *gorm.DB) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(models.Db)
	if err != nil {
		fmt.Println(err)
	}

	text :=
		`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
	m, _ := model.NewModelFromString(text)

	e, _ := casbin.NewEnforcer(m, a)
	return e
}

//主动更新全局的enforce对象
func loadPolicy() {
	en := GetEnforceObj()
	defer PutEnforceObj(en)
	en.LoadPolicy()
}

func GetEnforceObj() *casbin.Enforcer {
	en := syncPool.Get().(*casbin.Enforcer)
	return en
}

func PutEnforceObj(en *casbin.Enforcer) {
	en.LoadPolicy()
	syncPool.Put(en)
}

//添加规则
func AddPolice(rules [][]string) (bool, error) {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	ok, err := en.AddPolicies(rules)
	return ok, err
}

//删除规则
func RemovePolices(rules [][]string) (bool, error) {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	ok, err := en.RemovePolicies(rules)
	return ok, err
}

//获取全部分组
func GetGroups() []string {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	return en.GetAllNamedSubjects("p")
}

//获取组中是否具有该权限
func HasObjByGroup(groupname string, obj string, act string) bool {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	return en.HasPolicy(groupname, obj, act)
}

//将用户添加至指定分组
func AddGroup(ptype string, sub string, group string) (bool, error) {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	ok, err := en.AddNamedGroupingPolicy(ptype, sub, group)
	return ok, err
}

//批量添加用户权限
func AddGroups(ptype string, rules [][]string, txs ...*gorm.DB) (bool, error) {
	var en *casbin.Enforcer
	if len(txs) == 0 {
		en = GetEnforceObj()
		defer PutEnforceObj(en)
	} else {
		en = newEnforceObj(txs[0])
		defer loadPolicy()
	}

	ok, err := en.AddNamedGroupingPolicies(ptype, rules)
	return ok, err
}

//通过用户获取所在组
func GetGroupByUser(username string) ([]string, error) {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	privs, err := en.GetRolesForUser(username)
	return privs, err
}

//验证权限
func Check(sub string, obj string, act string) (bool, error) {

	if sub == "admin" {
		return true, nil
	}

	en := GetEnforceObj()
	defer PutEnforceObj(en)
	act = strings.ToLower(act)

	ok, err := en.Enforce(sub, obj, act)

	return ok, err
}

//删除权限
func DelGroups(ptype string, rules [][]string) (ok bool, err error) {
	en := GetEnforceObj()
	defer PutEnforceObj(en)
	ok, err = en.RemoveNamedPolicies(ptype, rules)
	return
}

//根据用户组获取权限
func GetPoliceByGroup(group string) [][]string {
	en := GetEnforceObj()
	defer PutEnforceObj(en)

	return en.GetFilteredPolicy(0, group)
}

//更新用户所属用户组
func UpdateGroups(username string, old []string, new []string, tx *gorm.DB) (ok bool, err error) {
	en := newEnforceObj(tx)

	add, incre := gstrings.CompareSlice(old, new)

	//添加新权限
	addLen := len(add)
	if addLen > 0 {
		addGropus := make([][]string, addLen)
		for addk, addv := range add {
			addGropus[addk] = []string{
				username,
				addv,
			}
		}

		ok, err = en.AddNamedGroupingPolicies("g", addGropus)
		if err != nil || !ok {
			return
		}
	}

	//删除旧权限
	increLen := len(incre)
	if increLen > 0 {
		increGropus := make([][]string, increLen)
		for increk, increv := range incre {
			increGropus[increk] = []string{
				username,
				increv,
			}
		}
		ok, err = en.RemoveNamedGroupingPolicies("g", increGropus)
		if err != nil || !ok {
			return
		}
	}

	return
}

//更新指定用户组的权限
func UpdatePolices(groupname string, old []string, new []string, tx *gorm.DB) (ok bool, err error) {
	addPolice, increPolice := gstrings.CompareSlice(old, new)
	en := newEnforceObj(tx)
	defer loadPolicy()

	addLen := len(addPolice)
	if addLen > 0 {
		var groupPrivs [][]string
		for _, v := range addPolice {
			privSlice := strings.Split(v, ":")
			groupPrivs = append(groupPrivs, []string{
				groupname,
				privSlice[0],
				privSlice[1],
			})
		}
		ok, err = en.AddPolicies(groupPrivs)
		if !ok || err != nil {
			return
		}
	}

	increLen := len(increPolice)
	if increLen > 0 {
		var groupIncrePrivs [][]string
		for _, v := range increPolice {
			privSlice := strings.Split(v, ":")
			groupIncrePrivs = append(groupIncrePrivs, []string{
				groupname,
				privSlice[0],
				privSlice[1],
			})
		}
		ok, err = en.RemovePolicies(groupIncrePrivs)
		if !ok || err != nil {
			return
		}
	}

	return
}

func GetRols() []string {
	en := GetEnforceObj()
	defer PutEnforceObj(en)
	roles := en.GetAllRoles()
	return roles
}
