package models

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
	// "github.com/pkg/errors"
	"github.com/astaxie/beego/logs"
)

type Menu struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	Link     string    `json:"link"`
	ParentId string    `json:"parent_id"`
	IsGroup  int       `json:"is_group"`
	IsSub    int       `json:"is_sub"`
	Status   int       `json:"status"`
	Sort     int       `json:"sort"`
	Chilren  []*Menu   `json:"chilren" xorm:"-"`
	Sub      []*Menu   `json:"sub" xorm:"-"`
	Created  int64     `json:"created" xorm:"created"`
	Updated  int64     `json:"updated" xorm:"updated"`
	Deleted  time.Time `json:"deleted" xorm:"deleted"`
}

type MenuWithPName struct {
	Menu       `xorm:"extends"`
	ParentName string `json:"parent_name" xorm:"parent_name"`
}

// 手动设置表名
func (m *Menu) TableName() string {
	return "menu"
}

type Menus []MenuWithPName
type MenusPaged struct {
	Menus `json:"menus"`
	Paged
}

func GetMenuRoot(parent_id string, r int) (Menu, error) {

	s := E.NewSession()
	defer s.Close()

	m := Menu{}
	logs.Error(r)
	if r > 5 {
		return m, nil
	}
	_, err := s.Where("id = ?", parent_id).Get(&m)
	if err != nil {
		return m, err
	} else {
		if m.ParentId != "" {
			m, err = GetMenuRoot(m.ParentId, r+1)
		}
	}

	return m, err
}

// 获取树状结构
func GetMenusTree(parent_id string, ids []string) ([]*Menu, error) {
	tree := []*Menu{}
	s := E.NewSession()
	defer s.Close()

	s.Where("parent_id = ?", parent_id)

	if len(ids) > 0 {
		if ids[0] != "*" {
			s.In("id", ids)
		}
	}
	err := s.Asc("sort").Desc("created").Find(&tree)
	if err != nil {
		return tree, nil
	}
	if len(tree) > 0 {
		for _, menu := range tree {
			s_tree, _ := GetMenusTree(menu.Id, ids)

			for _, s := range s_tree {
				if s.IsSub > 0 {
					menu.Sub = append(menu.Sub, s)
				} else {
					menu.Chilren = append(menu.Chilren, s)
				}
			}

		}
	}

	return tree, err
}

type MenuOption struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// 获取全部列表
func OptionList() ([]MenuOption, error) {
	options := []MenuOption{}
	err := E.Table("menu").Where("status = 0").Select("id,title").Asc("sort").Desc("created").Find(&options)
	return options, err
}

// 分页列表
func (m *Menu) Page(params url.Values) (MenusPaged, error) {
	menus := Menus{}
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
	err := E.Select("menu.*,(select title from menu as m where m.id = menu.parent_id) as parent_name").
		Limit(pageSize, (page-1)*pageSize).Asc("menu.sort").Desc("menu.created").
		Find(&menus)
	menusPaged := MenusPaged{menus, Paged{total, pageSize, page}}
	return menusPaged, err
}

//新增
func (m *Menu) Add(params url.Values) error {
	status, _ := strconv.Atoi(params.Get("status"))
	is_group, _ := strconv.Atoi(params.Get("is_group"))
	is_sub, _ := strconv.Atoi(params.Get("is_sub"))
	sort, _ := strconv.Atoi(params.Get("sort"))
	// parant_id == id 时递归查询错误
	// TODO 向上查询,parent_id 不能和ID形成环
	parent_id := params.Get("parent_id")
	root, _ := GetMenuRoot(parent_id,0)
	if root.Id == "" {
		libs.Logger.Error("parent_id 形成闭环")
		// return errors.New("parent_id 形成闭环")
		parent_id = ""
	}
	//插入帐号信息
	m = &Menu{
		Id:       uuid.Rand().Raw(),
		Title:    params.Get("title"),
		Icon:     params.Get("icon"),
		Link:     params.Get("link"),
		ParentId: parent_id,
		IsGroup:  is_group,
		IsSub:    is_sub,
		Sort:     sort,
		Status:   status,
	}
	_, err := E.Insert(m)
	if err != nil {
		return err
	}

	return err
}

// 编辑
func (m *Menu) Edit(params url.Values) error {

	status, _ := strconv.Atoi(params.Get("status"))
	is_group, _ := strconv.Atoi(params.Get("is_group"))
	is_sub, _ := strconv.Atoi(params.Get("is_sub"))
	sort, _ := strconv.Atoi(params.Get("sort"))
	// parant_id == id 时递归查询错误
	// TODO 向上查询,parent_id 不能和ID形成环
	parent_id := params.Get("parent_id")
	root, _ := GetMenuRoot(parent_id,0)
	logs.Error(root)
	if (params.Get("id") == parent_id) || root.Id == "" || root.Id == params.Get("id")  {
		libs.Logger.Error("parent_id 形成闭环")
		// return errors.New("parent_id 形成闭环")
		parent_id = ""
	}
	menu := &Menu{
		Title:    params.Get("title"),
		Icon:     params.Get("icon"),
		Link:     params.Get("link"),
		ParentId: parent_id,
		IsGroup:  is_group,
		IsSub:    is_sub,
		Sort:     sort,
		Status:   status,
	}

	//更新字段
	cols := []string{"title", "icon", "link", "parent_id", "is_group", "is_sub", "sort", "status"}
	_, err := E.Where("id = ?", params.Get("id")).Cols(cols...).Update(menu)
	if err != nil {
		return err
	} else {
		//补全接口未修改字段
		menu.Id = params.Get("id")
		menu.Created, _ = strconv.ParseInt(params.Get("created"), 10, 64)
		m = menu
	}

	return err
}

// 删除
func (m *Menu) Delete(params url.Values) error {
	//更新字段
	_, err := E.Where("id = ?", params.Get("id")).Delete(m)
	return err
}
