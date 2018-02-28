package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type ClustersController struct {
	controllers.BaseController
}

// 列表
// @Title Get Clusters List
// @Description  List Clusters
// @router / [post,get]
func (this *ClustersController) List() {
	var r controllers.Returned
	input := this.Input()

	cluster := models.Cluster{}
	data, err := cluster.Page(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 新增
// @Title Add a Region
// @Description  Add Region
// @router /add [post]
func (this *ClustersController) Add() {
	var r controllers.Returned

	input := this.Input()

	cluster := models.Cluster{}
	err := cluster.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = cluster
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a Region
// @Description  Edit Region
// @router /edit [post]
func (this *ClustersController) Edit() {
	var r controllers.Returned

	input := this.Input()

	cluster := models.Cluster{}
	err := cluster.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = cluster
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a Region
// @Description  Delete Region
// @router /delete [post]
func (this *ClustersController) Delete() {
	var r controllers.Returned

	input := this.Input()

	cluster := models.Cluster{}
	err := cluster.Delete(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = cluster
	}

	this.Data["json"] = r
	this.ServeJSON()
}

