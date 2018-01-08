package models

import (
	"net/url"
	"strings"

	"github.com/casbin/casbin"
)

func init() {
}

type Rule struct {
	Id      string `json:"id"`
	PType   string `json:"p_type" xorm:"'p_type'"`
	V0      string `json:"v0" xorm:"'v0'"`
	V1      string `json:"v1" xorm:"'v1'"`
	V2      string `json:"v2" xorm:"'v2'"`
	V3      string `json:"v3" xorm:"'v3'"`
	V4      string `json:"v4" xorm:"'v4'"`
	V5      string `json:"v5" xorm:"'v5'"`
	Created int64  `json:"created" xorm:"created"`
	Updated int64  `json:"updated" xorm:"updated"`
}
type Rules []Rule

// 手动设置表名
func (m *Rule) TableName() string {
	return "rule"
}

// 获取全部列表
func (m *Rule) List(params url.Values) (Rules, error) {
	rules := Rules{}

	s := E.NewSession()
	defer s.Close()

	s.Where("p_type = ? ",  params.Get("p_type"))
	if params.Get("v0") != "" {
		s.Where("v0 = ? ", params.Get("v0"))
	}
	if params.Get("v1") != "" {
		s.Where("v1 = ? ", params.Get("v1"))
	}

	err := s.Find(&rules)
	return rules, err
}

// 给角色设置用户
func (m *Rule) SetRoleUsers(params url.Values) (Rules, error) {
	rules := Rules{}
	e := casbin.NewEnforcer("conf/rbac.conf", Ca)
	err := e.LoadPolicy()
	role := params.Get("roles")

	cleansIds := params.Get("cleans")
	cleans := strings.Split(cleansIds, ",")
	for _, clean := range cleans {
		e.RemoveGroupingPolicy(clean, role)
	}

	usersIds := params.Get("users")
	if usersIds != "" {
		users := strings.Split(usersIds, ",")
		for _, user := range users {
			e.AddGroupingPolicy(user, role)
		}
	}
	E.Where("p_type = ? and v1 = ?", "g", role).Find(&rules)
	return rules, err
}

// 给用户设置角色
func (m *Rule) SetUserRoles(params url.Values) (Rules, error) {
	rules := Rules{}
	e := casbin.NewEnforcer("conf/rbac.conf", Ca)
	err := e.LoadPolicy()

	user := params.Get("users")

	cleansIds := params.Get("cleans")
	cleans := strings.Split(cleansIds, ",")
	for _, clean := range cleans {
		e.RemoveGroupingPolicy(user, clean)
	}

	rolesIds := params.Get("roles")
	if rolesIds != "" {
		roles := strings.Split(rolesIds, ",")
		for _, role := range roles {
			e.AddGroupingPolicy(user, role)
		}
	}

	E.Where("p_type = ? and v0 = ?", "g", user).Find(&rules)
	return rules, err
}

// 给角色设置菜单
func (m *Rule) SetRoleMenus(params url.Values) (Rules, error) {
	rules := Rules{}
	e := casbin.NewEnforcer("conf/rbac.conf", Ca)
	err := e.LoadPolicy()

	role := params.Get("roles")
	e.RemoveFilteredPolicy(0,role)
	menusIds := params.Get("menus")

	if menusIds != "" {
		menus := strings.Split(menusIds, ",")
		for _, menu := range menus {
			e.AddPolicy(role, menu)
		}
	}
	E.Where("p_type = ? and v0 = ?", "p", role).Find(&rules)
	return rules, err
}
