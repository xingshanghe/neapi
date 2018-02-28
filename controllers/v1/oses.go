package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type OsesController struct {
	controllers.BaseController
}

// 列表
// @Title Get Reions List
// @Description  List Roles
// @router / [post,get]
func (this *OsesController) List() {
	var r controllers.Returned
	input := this.Input()

	os := models.Os{}
	data, err := os.Page(input)
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
func (this *OsesController) Add() {
	var r controllers.Returned

	input := this.Input()

	os := models.Os{}
	err := os.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = os
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a Region
// @Description  Edit Region
// @router /edit [post]
func (this *OsesController) Edit() {
	var r controllers.Returned

	input := this.Input()

	os := models.Os{}
	err := os.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = os
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a Region
// @Description  Delete Region
// @router /delete [post]
func (this *OsesController) Delete() {
	var r controllers.Returned

	input := this.Input()

	os := models.Os{}
	err := os.Delete(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = os
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 树状结构菜单列表
// @Title Get Region Options
// @Description  Get Region Options
// @router /options [post,get]
func (this *OsesController) Options() {
	var r controllers.Returned
	data, err := models.OsOptionList()
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}