package models

import (
	"github.com/xingshanghe/neapi/libs"
	"net/url"
	"strconv"
)

func init() {
}

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type Detail struct {
	Id        int    `json:"id"`
	AccountId int    `json:"account_id" xorm:"index"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

type User struct {
	Account `xorm:"extends"`
	Detail  `xorm:"extends"`
}

func (m *User) TableName() string {
	return "account"
}

// 手动设置表名
func (m *Detail) TableName() string {
	return "account_detail"
}

type Users []User
type UsersPaged struct {
	Users `json:"users"`
	Paged
}

// 获取全部列表
func (m *User) List() (Users, error) {
	users := Users{}

	err := E.Join("LEFT OUTER", []string{m.Detail.TableName(), "d"}, "account.id = d.account_id").
		Find(&users)

	return users, err
}

// 分页列表
func (m *User) Page(params url.Values) (UsersPaged, error) {
	users := Users{}
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
	err := E.Join("INNER", []string{m.Detail.TableName(), "d"}, "account.id = d.account_id").
		Limit(pageSize, (page-1)*pageSize).
		Find(&users)
	userPaged := UsersPaged{users, Paged{total, pageSize, page}}
	return userPaged, err
}
