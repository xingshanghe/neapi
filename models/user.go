package models

import (
	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
	"net/url"
	"strconv"
	"time"
)

func init() {
}

type Account struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Status   int       `json:"status"`
	Deleted  time.Time `json:"deleted" xorm:"deleted"`
}

type Detail struct {
	Id        string    `json:"did"`
	AccountId string    `json:"account_id" xorm:"index"`
	Nickname  string    `json:"nickname"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	Birthday  string    `json:"birthday"`
	Created   int64     `json:"created" xorm:"created"`
	Updated   int64     `json:"updated" xorm:"updated"`
	Deleted   time.Time `json:"detail_deleted" xorm:"deleted"`
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

	err := E.Join("INNER", []string{m.Detail.TableName(), "d"}, "account.id = d.account_id").
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
		Limit(pageSize, (page-1)*pageSize).Desc("d.created").
		Find(&users)
	userPaged := UsersPaged{users, Paged{total, pageSize, page}}
	return userPaged, err
}

func (m *User) Add(params url.Values) error {
	s := E.NewSession()
	defer s.Close()

	err := s.Begin()
	if err != nil {
		return err
	}
	//插入帐号信息
	account := Account{
		Id:       uuid.Rand().Raw(),
		Username: params.Get("username"),
		Password: libs.GetRandomString(6),
		Email:    params.Get("email"),
		Phone:    params.Get("phone"),
		Status:   0,
	}
	_, err = s.Insert(&account)
	if err != nil {
		s.Rollback()
		return err
	}
	age, _ := strconv.Atoi(params.Get("age"))
	detail := Detail{
		Id:        uuid.Rand().Raw(),
		AccountId: account.Id,
		Nickname:  params.Get("nickname"),
		Gender:    params.Get("gender"),
		Age:       age,
		Address:   params.Get("address"),
		Birthday:  params.Get("birthday"),
	}
	_, err = s.Insert(&detail)
	if err != nil {
		s.Rollback()
		return err
	}
	err = s.Commit()
	if err != nil {
		return err
	} else {
		m.Account = account
		m.Detail = detail
	}

	return err
}

func (m *User) Edit(params url.Values) error {
	s := E.NewSession()
	defer s.Close()

	err := s.Begin()
	if err != nil {
		return err
	}

	account := Account{
		Username: params.Get("username"),
		Email:    params.Get("email"),
		Phone:    params.Get("phone"),
	}

	//更新字段
	_, err = E.Where("id = ?", params.Get("id")).Update(&account)
	if err != nil {
		s.Rollback()
		return err
	}
	age, _ := strconv.Atoi(params.Get("age"))
	detail := Detail{
		Nickname: params.Get("nickname"),
		Gender:   params.Get("gender"),
		Age:      age,
		Address:  params.Get("address"),
		Birthday: params.Get("birthday"),
	}
	//更新字段
	_, err = E.Where("id = ?", params.Get("did")).Update(&detail)

	if err != nil {
		s.Rollback()
		return err
	}

	err = s.Commit()
	if err != nil {
		return err
	} else {
		//补全接口未修改字段,不补全可以看出那些字段被修改过
		account.Id = params.Get("id")
		account.Password = params.Get("password")
		account.Status, _ = strconv.Atoi(params.Get("status"))
		detail.Id = params.Get("did")
		detail.AccountId = params.Get("account_id")
		detail.Created, _ = strconv.ParseInt(params.Get("created"), 10, 64)

		m.Account = account
		m.Detail = detail
	}

	return err
}

func (m *User) Delete(params url.Values) error {
	s := E.NewSession()
	defer s.Close()

	err := s.Begin()
	if err != nil {
		return err
	}

	account := Account{}

	//更新字段
	_, err = E.Where("id = ?", params.Get("id")).Delete(&account)
	if err != nil {
		s.Rollback()
		return err
	}

	detail := Detail{}
	//更新字段
	_, err = E.Where("id = ?", params.Get("did")).Delete(&detail)

	if err != nil {
		s.Rollback()
		return err
	}

	err = s.Commit()
	if err != nil {
		return err
	}

	return err
}

// 切换用户状态
func (m *User) ToggleStatus(params url.Values) error {
	s := E.NewSession()
	defer s.Close()

	err := s.Begin()
	if err != nil {
		return err
	}

	statusOld, _ := strconv.Atoi(params.Get("status"))
	var status int
	if statusOld == 0 {
		status = 1
	}
	if statusOld == 1 {
		status = 0
	}
	account := Account{
		Status: status,
	}

	//更新字段
	_, err = E.Where("id = ?", params.Get("id")).Cols("status").Update(&account)
	if err != nil {
		s.Rollback()
		return err
	}
	detail := Detail{}
	//更新字段
	_, err = E.Where("id = ?", params.Get("did")).Update(&detail)

	if err != nil {
		s.Rollback()
		return err
	}

	err = s.Commit()
	if err != nil {
		return err
	} else {
		m.Account = account
		m.Detail = detail
	}

	return err
}

// 重置密码
func (m *User) ResetPwd(params url.Values) error {
	s := E.NewSession()
	defer s.Close()

	err := s.Begin()
	if err != nil {
		return err
	}

	account := Account{
		Password: libs.GetRandomString(6),
	}

	//更新字段
	_, err = E.Where("id = ?", params.Get("id")).Update(&account)
	if err != nil {
		s.Rollback()
		return err
	}
	detail := Detail{}
	//更新字段
	_, err = E.Where("id = ?", params.Get("did")).Update(&detail)

	if err != nil {
		s.Rollback()
		return err
	}

	err = s.Commit()
	if err != nil {
		return err
	} else {
		m.Account = account
		m.Detail = detail
	}

	return err
}
