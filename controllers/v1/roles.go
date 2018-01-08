package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type RolesController struct {
	controllers.BaseController
}

// 列表
// @Title Get User Role List
// @Description  List Roles
// @router / [post,get]
func (this *RolesController) List() {
	var r controllers.Returned
	input := this.Input()

	role := models.Role{}
	data, err := role.Page(input)
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
func (this *RolesController) Add() {
	var r controllers.Returned

	input := this.Input()

	role := models.Role{}
	err := role.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = role
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a User
// @Description  Edit User
// @router /edit [post]
func (this *RolesController) Edit() {
	var r controllers.Returned

	input := this.Input()

	role := models.Role{}
	err := role.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = role
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a User
// @Description  Delete User
// @router /delete [post]
func (this *RolesController) Delete() {
	var r controllers.Returned

	input := this.Input()

	role := models.Role{}
	err := role.Delete(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = role
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 为角色添加用户
// @Title set Roles to a User
// @Description  set Roles to a User
// @router /setUsers [post]
func (this *RolesController) SetUsers() {
	var r controllers.Returned
	input := this.Input()
	rule := models.Rule{}
	rules, err := rule.SetRoleUsers(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = rules
	}
	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title set rules
// @Description  set rules
// @router /setMenus [post]
func (this *RolesController) SetMenus() {
	var r controllers.Returned

	input := this.Input()

	rule := models.Rule{}
	rules, err := rule.SetRoleMenus(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = rules
	}

	this.Data["json"] = r
	this.ServeJSON()
}
