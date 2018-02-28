package models

import (
	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
	"net/url"
	"strconv"
	"time"
)

type Role struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Created     int64     `json:"created" xorm:"created"`
	Updated     int64     `json:"updated" xorm:"updated"`
	Deleted     time.Time `json:"deleted" xorm:"deleted"`
}

// 手动设置表名
func (m *Role) TableName() string {
	return "role"
}

type Roles []Role
type RolesPaged struct {
	Roles `json:"roles"`
	Paged
}

// 获取全部列表
func (m *Role) List() (Roles, error) {
	roles := Roles{}
	err := E.Find(&roles)
	return roles, err
}

// 分页列表
func (m *Role) Page(params url.Values) (RolesPaged, error) {
	roles := Roles{}
	//分页处理
	pageSize, e1 := strconv.Atoi(params.Get("pageSize"))
	if e1 != nil {
		pageSize = libs.PageSize
	}
	page, e1 := strconv.Atoi(params.Get("page"))
	if e1 != nil {
		page = libs.PageNo
	} else {
		if page < 1 {
			page = 1
		}
	}
	total, _ := E.Count(m)
	err := E.Limit(pageSize, (page-1)*pageSize).Desc("created").
		Find(&roles)
	rolesPaged := RolesPaged{roles, Paged{total, pageSize, page}}
	return rolesPaged, err
}

//新增
func (m *Role) Add(params url.Values) error {
	status, _ := strconv.Atoi(params.Get("status"))
	//插入帐号信息
	m = &Role{
		Id:          uuid.Rand().Raw(),
		Name:        params.Get("name"),
		Code:        params.Get("code"),
		Description: params.Get("description"),
		Status:      status,
	}
	_, err := E.Insert(m)
	if err != nil {
		return err
	}

	return err
}

// 编辑
func (m *Role) Edit(params url.Values) error {

	status, _ := strconv.Atoi(params.Get("status"))
	role := &Role{
		Name:        params.Get("name"),
		Code:        params.Get("code"),
		Description: params.Get("description"),
		Status:      status,
	}

	//更新字段
	cols := []string{"name", "code", "status", "description"}
	_, err := E.Where("id = ?", params.Get("id")).Cols(cols...).Update(role)
	if err != nil {
		return err
	} else {
		//补全接口未修改字段
		role.Id = params.Get("id")
		role.Created, _ = strconv.ParseInt(params.Get("created"), 10, 64)
		m = role
	}

	return err
}

// 删除
func (m *Role) Delete(params url.Values) error {
	//更新字段
	_, err := E.Where("id = ?", params.Get("id")).Delete(m)
	return err
}
