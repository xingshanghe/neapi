package models

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
)

type NodePrimary struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	RegionId    string `json:"region_id" xorm:"index"`
	OsId        string `json:"os_id" xorm:"index"`
	ClusterId   string `json:"cluster_id" xorm:"cluster_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	Status      int    `json:"status"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
	// CreatedBy   string    `json:"created_by" xorm:"index"`
	Created int64     `json:"created" xorm:"created"`
	Updated int64     `json:"updated" xorm:"updated"`
	Deleted time.Time `json:"deleted" xorm:"deleted"`
}

type Node struct {
	NodePrimary `xorm:"extends" json:"node"`
	Region      `xorm:"extends" json:"region"`
	Os          `xorm:"extends" json:"os"`
}

// 手动设置表名
func (m *Node) TableName() string {
	return "node"
}
// 手动设置表名
func (m *NodePrimary) TableName() string {
	return "node"
}

type Nodes []Node
type NodesPaged struct {
	Nodes `json:"nodes"`
	Paged
}

// 获取全部列表
func (m *Node) List(params url.Values) (Nodes, error) {
	nodes := Nodes{}
	err := E.Find(&nodes)
	return nodes, err
}

// 分页列表
func (m *Node) Page(params url.Values) (NodesPaged, error) {
	nodes := Nodes{}
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
	s := E.NewSession()
	defer s.Close()
	if params.Get("idle") == "1" {
		s.Where("cluster_id = ''")
	}
	total, _ := s.Count(m)

	err := s.Select("*").
		Join("LEFT", "region", "node.region_id = region.id").
		//Join("INNER", "cluster", "node.cluster_id = cluster.id").
		Join("LEFT", "os", "node.os_id = os.id").
		Limit(pageSize, (page-1)*pageSize).Asc("node.sort").Desc("node.created").
		Find(&nodes)
	nodesPaged := NodesPaged{nodes, Paged{total, pageSize, page}}
	return nodesPaged, err
}

//新增
func (m *Node) Add(params url.Values) error {

	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))
	//插入帐号信息
	np := &NodePrimary{
		Id:          uuid.Rand().Raw(),
		Name:        params.Get("name"),
		Alias:       params.Get("alias"),
		Ip:          params.Get("ip"),
		Port:        params.Get("port"),
		Username:    params.Get("username"),
		Password:    params.Get("password"),
		RegionId:    params.Get("region_id"),
		OsId:        params.Get("os_id"),
		Status:      status,
		Sort:        sort,
		Description: params.Get("description"),
	}
	_, err := E.Insert(np)
	if err != nil {
		return err
	} else {
		//m.NodePrimary.Id = np.Id
		nodes := Nodes{}
		err := E.Select("*").
			Join("LEFT", "region", "node.region_id = region.id").
			//Join("INNER", "cluster", "node.cluster_id = cluster.id").
			Join("LEFT", "os", "node.os_id = os.id").
			Where("node.id = ?", np.Id).
			Limit(1).
			Find(&nodes)
		if err != nil {
			return err
		} else {
			if len(nodes) > 0 {
				*m = nodes[0]
			}
		}
	}

	return err
}

// 编辑
func (m *Node) Edit(params url.Values) error {

	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))

	m.NodePrimary.Name = params.Get("name")
	m.NodePrimary.Alias = params.Get("alias")
	m.NodePrimary.Ip = params.Get("ip")
	m.NodePrimary.Port = params.Get("port")
	m.NodePrimary.Username = params.Get("username")
	m.NodePrimary.Password = params.Get("password")
	m.NodePrimary.RegionId = params.Get("region_id")
	m.NodePrimary.OsId = params.Get("os_id")
	m.NodePrimary.Status = status
	m.NodePrimary.Sort = sort
	m.NodePrimary.Description = params.Get("description")

	//更新字段
	cols := []string{"name", "alias", "ip", "port", "username", "password", "region_id", "os_id", "status", "sort", "description"}
	_, err := E.Where("id = ?", params.Get("id")).Cols(cols...).Update(&m.NodePrimary)
	if err != nil {
		return err
	}

	return err
}

// 删除
func (m *Node) Delete(params url.Values) error {
	//更新字段
	_, err := E.Where("id = ?", params.Get("id")).Delete(m)
	return err
}
