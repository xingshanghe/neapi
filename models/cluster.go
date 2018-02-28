package models

import (
	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/libs/uuid"
	"net/url"
	"strconv"
	"time"
)

type Cluster struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	IpCidr      string    `json:"ip_cidr"`
	StoreDir    string    `json:"store_dir"`
	Status      int       `json:"status"`
	Sort        int       `json:"sort"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by" xorm:"index"`
	Created     int64     `json:"created" xorm:"created"`
	Updated     int64     `json:"updated" xorm:"updated"`
	Deleted     time.Time `json:"deleted" xorm:"deleted"`
}

// 手动设置表名
func (m *Cluster) TableName() string {
	return "cluster"
}

type Clusters []Cluster
type ClustersPaged struct {
	Clusters `json:"clusters"`
	Paged
}

// 获取全部列表
func (m *Cluster) List(params url.Values) (Clusters, error) {
	clusters := Clusters{}
	err := E.Find(&clusters)
	return clusters, err
}

// 分页列表
func (m *Cluster) Page(params url.Values) (ClustersPaged, error) {
	clusters := Clusters{}
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
		Find(&clusters)
	clustersPaged := ClustersPaged{clusters, Paged{total, pageSize, page}}
	return clustersPaged, err
}

//新增
func (m *Cluster) Add(params url.Values) error {
	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))
	//插入帐号信息
	m = &Cluster{
		Id:          uuid.Rand().Raw(),
		Name:        params.Get("name"),
		Alias:       params.Get("alias"),
		IpCidr:      params.Get("ip_cidr"),
		StoreDir:    params.Get("store_dir"),
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
func (m *Cluster) Edit(params url.Values) error {

	status, _ := strconv.Atoi(params.Get("status"))
	sort, _ := strconv.Atoi(params.Get("sort"))

	cluster := &Cluster{
		Name:        params.Get("name"),
		Alias:       params.Get("alias"),
		IpCidr:      params.Get("ip_cidr"),
		StoreDir:    params.Get("store_dir"),
		Status:      status,
		Sort:        sort,
		Description: params.Get("description"),
	}

	//更新字段
	cols := []string{"name", "alias", "ip_cidr", "store_dir", "status", "sort", "description"}
	_, err := E.Where("id = ?", params.Get("id")).Cols(cols...).Update(cluster)
	if err != nil {
		return err
	} else {
		//补全接口未修改字段
		cluster.Id = params.Get("id")
		cluster.Created, _ = strconv.ParseInt(params.Get("created"), 10, 64)
		m = cluster
	}

	return err
}

// 删除
func (m *Cluster) Delete(params url.Values) error {
	//更新字段
	_, err := E.Where("id = ?", params.Get("id")).Delete(m)
	return err
}
