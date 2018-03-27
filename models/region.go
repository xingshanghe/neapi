package models

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
)

type Region struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Status      int       `json:"status"`
	Sort        int       `json:"sort"`
	Description string    `json:"description"`
	Created     int64     `json:"created" xorm:"created"`
	Updated     int64     `json:"updated" xorm:"updated"`
	Deleted     time.Time `json:"deleted" xorm:"deleted"`
}

// 手动设置表名
func (m *Region) TableName() string {
	return "region"
}

type Regions []Region
type RegionsPaged struct {
	Regions `json:"regions"`
	Paged
}
type RegionOption struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// 获取全部列表
func RegionOptionList() ([]RegionOption, error) {
	options := []RegionOption{}
	err := E.Table("region").Select("id,name,code").Where("status = 0").Asc("sort").Desc("created").Find(&options)
	return options, err
}

// 获取全部列表
func (m *Region) List(params url.Values) (Regions, error) {
	regions := Regions{}
	err := E.Find(&regions)
	return regions, err
}

// 分页列表
func (m *Region) Page(params url.Values) (RegionsPaged, error) {
	regions := Regions{}
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
		Find(&regions)
	regionsPaged := RegionsPaged{regions, Paged{total, pageSize, page}}
	return regionsPaged, err
}

//新增
func (m *Region) Add(params url.Values) error {
	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))
	//插入帐号信息
	m = &Region{
		Id:          uuid.Rand().Raw(),
		Name:        params.Get("name"),
		Code:        params.Get("code"),
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
func (m *Region) Edit(params url.Values) error {

	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))

	m.Name = params.Get("name")
	m.Code = params.Get("code")
	m.Status = status
	m.Sort = sort
	m.Description = params.Get("description")

	//更新字段
	cols := []string{"name", "code", "status", "sort", "description"}
	_, err := E.Where("id = ?", m.Id).Cols(cols...).Update(m)
	if err != nil {
		return err
	}

	return err
}

// 删除
func (m *Region) Delete(params url.Values) error {
	//更新字段
	_, err := E.Where("id = ?", m.Id).Delete(m)
	return err
}
