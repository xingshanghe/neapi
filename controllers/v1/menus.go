package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type MenusController struct {
	controllers.BaseController
}

// 列表
// @Title Get User Role List
// @Description  List Roles
// @router / [post,get]
func (this *MenusController) List() {
	var r controllers.Returned
	input := this.Input()

	menu := models.Menu{}
	data, err := menu.Page(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 树状结构菜单列表
// @Title Get Menu Tree
// @Description  Get Menu Tree
// @router /options [post,get]
func (this *MenusController) Options() {
	var r controllers.Returned
	//根据角色查询 菜单
	//ids := models.GetMenusIdsByRoles(roles)
	data, err := models.OptionList()
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 树状结构菜单列表,用于系统，系统设置时使用
// @Title Get Menu Tree
// @Description  Get Menu Tree
// @router /tree [post,get]
func (this *MenusController) Tree() {
	var r controllers.Returned
	//根据角色查询 菜单
	//ids := models.GetMenusIdsByRoles(roles)
	data, err := models.GetMenusTree("", []string{}, false)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 树状结构菜单列表,用于系统，系统设置时使用
// @Title Get Menu Tree
// @Description  Get Menu Tree
// @router /layouts/tree [post,get]
func (this *MenusController) RoleTree() {
	var r controllers.Returned
	//根据角色查询 菜单
	//input := this.Input()
	data, err := models.GetMenusTree("", []string{}, false)
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
// @Title Add a User
// @Description  Add User
// @router /add [post]
func (this *MenusController) Add() {
	var r controllers.Returned

	input := this.Input()

	menu := models.Menu{}
	err := menu.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = menu
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a User
// @Description  Edit User
// @router /edit [post]
func (this *MenusController) Edit() {
	var r controllers.Returned

	input := this.Input()

	menu := models.Menu{}
	err := menu.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = menu
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a User
// @Description  Delete User
// @router /delete [post]
func (this *MenusController) Delete() {
	var r controllers.Returned

	input := this.Input()

	menu := models.Menu{}
	err := menu.Delete(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = menu
	}

	this.Data["json"] = r
	this.ServeJSON()
}
