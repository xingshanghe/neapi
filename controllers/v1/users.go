package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/models"
	"net/url"
)

type UsersController struct {
	controllers.BaseController
}

// 测试方法
// @Title 测试方法
// @Description  测试方法
// @router /test [get]
func (this *UsersController) Test() {
	var r controllers.Returned

	r.Data = libs.MD5("xiNgsHangHewaNSui")

	this.Data["json"] = r
	this.ServeJSON()
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

// 获取用户详情
// @Title Get User Account List
// @Description  List Users
// @router /:username [post,get]
func (this *UsersController) Detail()  {
	var r controllers.Returned

	username := this.Ctx.Input.Param(":username")
	account := models.Account{
		Username: username,
	}
	hasThisAccount, err := models.E.Get(&account)
	if err == nil {
		if hasThisAccount {
			detail := models.Detail{
				AccountId: account.Id,
			}
			models.E.Get(&detail)
			user := models.User{
				account,
				detail,
			}
			r.Data = user
		}
	}else{
		r.Code = 5000
		r.Msg = err.Error()
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑（管理员修改用户信息）
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

// 编辑（修改用户信息,完成后需要重新登录）
// @Title Edit a User
// @Description  Edit User
// @router /editSelf [post]
func (this *UsersController) EditSelf() {
	var r controllers.Returned

	input := this.Input()

	user := models.User{}
	err := user.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		var data struct {
			Token string      `json:"token"`
			User  models.User `json:"account"`
		}
		data.User = user
		roleIds := []string{}

		menuRule := models.MenuRule{}
		p := url.Values{}
		p.Set("p_type", "g")
		p.Set("v0", user.AccountId)
		menuRules, _ := menuRule.List(p)
		for _, mr := range menuRules {
			roleIds = append(roleIds, mr.V1)
		}

		var token string
		token, err = libs.CreateJwt(user, roleIds)
		data.Token = token
		r.Data = data
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

// 重置密码
// @Title ToggleStatus a User
// @Description  ToggleStatus User
// @router /editPwd [post]
func (this *UsersController) EditPwd() {
	var r controllers.Returned

	input := this.Input()

	password := input.Get("password")
	password0 := input.Get("password0")
	password1 := input.Get("password1")
	password2 := input.Get("password2")

	if password1 == password2 {

		if models.GetPassword(password0) == password {
			user := models.User{}
			input.Set("password", password1)
			err := user.ResetPwd(input)
			if err != nil {
				r.Code = 5000
				r.Msg = err.Error()
			} else {
				r.Data = user
			}
		} else {
			r.Code = 5000
			r.Msg = "原始密码不正确！"
		}
	} else {
		r.Code = 5000
		r.Msg = "两次密码不一致！"
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
