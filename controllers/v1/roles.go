package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
	"net/url"
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
	menuRule := models.MenuRule{}
	menuRules, err := menuRule.SetRoleUsers(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = menuRules
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

	menuRule := models.MenuRule{}
	menuRules, err := menuRule.SetRoleMenus(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = menuRules
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 树状结构菜单列表,用于系统，系统设置时使用
// @Title Get Menu Tree
// @Description  Get Menu Tree
// @router /tree [post,get]
func (this *RolesController) Tree() {
	var r controllers.Returned
	//根据角色查询 菜单
	input := this.Input()
	roleIds := input.Get("role_ids")
	menuIds := []string{}

	if roleIds != "" {
		menuRule := models.MenuRule{}
		p := url.Values{}
		p.Set("p_type", "p")
		p.Set("v0", roleIds)
		menuRules, _ := menuRule.List(p)

		for _, mr := range menuRules {
			menuIds = append(menuIds, mr.V1)
		}
	}

	//data, err := models.GetMenusTree("", menuIds, true)
	data, err := models.GetMenusTree("", menuIds, true)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}
