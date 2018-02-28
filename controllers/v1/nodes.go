package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
	"github.com/astaxie/beego/logs"
)

type NodesController struct {
	controllers.BaseController
}

// 空闲列表
// @Title Get Nodes List
// @Description  List Nodes
// @router / [post,get]
func (this *NodesController) List() {
	var r controllers.Returned
	input := this.Input()

	input.Set("idle","1")

	node := models.Node{}
	data, err := node.Page(input)
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
// @Title Add a Node
// @Description  Add Node
// @router /add [post]
func (this *NodesController) Add() {
	var r controllers.Returned

	input := this.Input()

	node := models.Node{}
	err := node.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		logs.Error(node)
		r.Data = node
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a Node
// @Description  Edit Node
// @router /edit [post]
func (this *NodesController) Edit() {
	var r controllers.Returned

	input := this.Input()

	node := models.Node{}
	err := node.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = node
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a Node
// @Description  Delete Node
// @router /delete [post]
func (this *NodesController) Delete() {
	var r controllers.Returned

	input := this.Input()

	node := models.Node{}
	err := node.Delete(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = node
	}

	this.Data["json"] = r
	this.ServeJSON()
}