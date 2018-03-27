package models

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
)

type Os struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Family      string    `json:"family"`
	Version     string    `json:"version"`
	Bit         int       `json:"bit"`
	Status      int       `json:"status"`
	Sort        int       `json:"sort"`
	Description string    `json:"description"`
	Created     int64     `json:"created" xorm:"created"`
	Updated     int64     `json:"updated" xorm:"updated"`
	Deleted     time.Time `json:"deleted" xorm:"deleted"`
}

// 手动设置表名
func (m *Os) TableName() string {
	return "os"
}

type Oses []Os
type OsesPaged struct {
	Oses `json:"oses"`
	Paged
}

type OsOption struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Family  string `json:"family"`
	Version string `json:"version"`
	Bit     int    `json:"bit"`
}

// 获取全部列表
func OsOptionList() ([]OsOption, error) {
	options := []OsOption{}
	//(`node`.`deleted` IS NULL OR `node`.`deleted`=?)
	err := E.Table("os").Select("id,name,family,version,bit").Where("status = 0").And("deleted is NULL or deleted = ?", "001-01-01 00:00:00").Asc("sort").Desc("created").Find(&options)
	return options, err
}

// 获取全部列表
func (m *Os) List(params url.Values) (Oses, error) {
	oss := Oses{}
	err := E.Find(&oss)
	return oss, err
}

// 分页列表
func (m *Os) Page(params url.Values) (OsesPaged, error) {
	oses := Oses{}
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
	err := E.Select("*").
		Limit(pageSize, (page-1)*pageSize).Asc("sort").Desc("created").
		Find(&oses)
	osesPaged := OsesPaged{oses, Paged{total, pageSize, page}}
	return osesPaged, err
}

//新增
func (m *Os) Add(params url.Values) error {
	bit, _ := strconv.Atoi(params.Get("bit"))
	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))
	//插入帐号信息
	m = &Os{
		Id:          uuid.Rand().Raw(),
		Name:        params.Get("name"),
		Family:      params.Get("family"),
		Version:     params.Get("version"),
		Bit:         bit,
		Status:      status,
		Sort:        sort,
		Description: params.Get("description"),
	}
	_, err := E.Insert(m)
	if err != nil {
		return err
	}

	return err
}

// 编辑
func (m *Os) Edit(params url.Values) error {
	bit, _ := strconv.Atoi(params.Get("bit"))
	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))

	m.Name = params.Get("name")
	m.Family = params.Get("family")
	m.Version = params.Get("version")
	m.Bit = bit
	m.Status = status
	m.Sort = sort
	m.Description = params.Get("description")

	//更新字段
	cols := []string{"name", "family", "version", "bit", "status", "sort", "description"}
	_, err := E.Where("id = ?", m.Id).Cols(cols...).Update(m)
	if err != nil {
		return err
	}

	return err
}

// 删除
func (m *Os) Delete(params url.Values) error {
	//更新字段
	_, err := E.Where("id = ?", m.Id).Delete(m)
	return err
}
