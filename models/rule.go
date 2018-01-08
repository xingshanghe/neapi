package models

import (
	"net/url"
	"strings"
)

func init() {
}

type ApiRule struct {
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
type ApiRules []ApiRule

// 手动设置表名
func (m *ApiRule) TableName() string {
	return "api_rule"
}

type MenuRule struct {
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
type MenuRules []MenuRule

// 手动设置表名
func (m *MenuRule) TableName() string {
	return "menu_rule"
}

// 获取全部列表
func (m *MenuRule) List(params url.Values) (MenuRules, error) {
	menuRules := MenuRules{}

	s := E.NewSession()
	defer s.Close()

	s.Where("p_type = ? ", params.Get("p_type"))
	if params.Get("v0") != "" {
		s.In("v0",strings.Split(params.Get("v0"), ","))
		//s.Where("v0 = ? ", params.Get("v0"))
	}
	if params.Get("v1") != "" {
		s.In("v1",strings.Split(params.Get("v1"), ","))
		//s.Where("v0 = ? ", params.Get("v0"))
		//s.Where("v1 = ? ", params.Get("v1"))
	}

	err := s.Find(&menuRules)
	return menuRules, err
}

// 给角色设置用户
func (m *MenuRule) SetRoleUsers(params url.Values) (MenuRules, error) {
	menuRules := MenuRules{}
	role := params.Get("roles")

	cleansIds := params.Get("cleans")
	cleans := strings.Split(cleansIds, ",")
	for _, clean := range cleans {
		Cme.RemoveGroupingPolicy(clean, role)
	}

	usersIds := params.Get("users")
	if usersIds != "" {
		users := strings.Split(usersIds, ",")
		for _, user := range users {
			Cme.AddGroupingPolicy(user, role)
		}
	}
	err := Cme.LoadPolicy()
	if err != nil {
		return menuRules, nil
	}
	err = E.Where("p_type = ? and v1 = ?", "g", role).Find(&menuRules)
	return menuRules, err
}

// 给用户设置角色
func (m *MenuRule) SetUserRoles(params url.Values) (MenuRules, error) {
	menuRules := MenuRules{}

	user := params.Get("users")

	cleansIds := params.Get("cleans")
	cleans := strings.Split(cleansIds, ",")
	for _, clean := range cleans {
		Cme.RemoveGroupingPolicy(user, clean)
	}

	rolesIds := params.Get("roles")
	if rolesIds != "" {
		roles := strings.Split(rolesIds, ",")
		for _, role := range roles {
			Cme.AddGroupingPolicy(user, role)
		}
	}

	err := Cme.LoadPolicy()
	if err != nil {
		return menuRules, nil
	}

	err = E.Where("p_type = ? and v0 = ?", "g", user).Find(&menuRules)
	return menuRules, err
}

// 给角色设置菜单
func (m *MenuRule) SetRoleMenus(params url.Values) (MenuRules, error) {
	menuRules := MenuRules{}

	role := params.Get("roles")
	//暂时没找到相应的高级方法，可以用数据库操作代替
	Cme.RemoveFilteredPolicy(0, role)
	menusIds := params.Get("menus")

	if menusIds != "" {
		menus := strings.Split(menusIds, ",")
		for _, menu := range menus {
			Cme.AddPolicy(role, menu)
		}
	}
	err := Cme.LoadPolicy()
	if err != nil {
		return menuRules, nil
	}
	err = E.Where("p_type = ? and v0 = ?", "p", role).Find(&menuRules)
	return menuRules, err
}

func (m *MenuRule) GetMenus(params url.Values) ([]string, error) {
	menuRules := MenuRules{}
	E.Where("p_type = ? and v0 = ?", "p", params.Get("role_id")).Find(&menuRules)
	//return menuRules, err
	return []string{}, nil
}
