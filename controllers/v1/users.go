package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type UsersController struct {
	controllers.BaseController
}

// 列表
// @Title Get User Account List
// @Description  List Users
// @router / [post,get]
func (this *UsersController) List() {
	var r controllers.Returned
	input := this.Input()

	user := models.User{}
	data, err := user.Page(input)
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
func (this *UsersController) Add() {
	var r controllers.Returned

	input := this.Input()

	user := models.User{}
	err := user.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = user
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a User
// @Description  Edit User
// @router /edit [post]
func (this *UsersController) Edit() {
	var r controllers.Returned

	input := this.Input()

	user := models.User{}
	err := user.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = user
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a User
// @Description  Delete User
// @router /delete [post]
func (this *UsersController) Delete() {
	var r controllers.Returned

	input := this.Input()

	if input.Get("username") != "admin" {
		user := models.User{}
		err := user.Delete(input)
		if err != nil {
			r.Code = 5000
			r.Msg = err.Error()
		} else {
			r.Data = user
		}
	} else {
		r.Code = 5001
		r.Msg = "系统保留帐号，禁止删除!"
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 切换用户状态
// @Title ToggleStatus a User
// @Description  ToggleStatus User
// @router /toggleStatus [post]
func (this *UsersController) ToggleStatus() {
	var r controllers.Returned

	input := this.Input()

	user := models.User{}
	err := user.ToggleStatus(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = user
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 重置密码
// @Title ToggleStatus a User
// @Description  ToggleStatus User
// @router /resetPwd [post]
func (this *UsersController) ResetPwd() {
	var r controllers.Returned

	input := this.Input()

	user := models.User{}
	err := user.ResetPwd(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = user
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 为用户添加角色
// @Title set Roles to a User
// @Description  set Roles to a User
// @router /setRoles [post]
func (this *UsersController) SetRoles() {
	var r controllers.Returned
	input := this.Input()
	menuRule := models.MenuRule{}
	menuRules, err := menuRule.SetUserRoles(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = menuRules
	}
	this.Data["json"] = r
	this.ServeJSON()
}
