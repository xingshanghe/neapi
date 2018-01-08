package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type RulesController struct {
	controllers.BaseController
}

// g列表
// @Title Get User Role List
// @Description  List Roles
// @router /g [post,get]
func (this *RulesController) GList() {
	var r controllers.Returned
	input := this.Input()

	menuRule := models.MenuRule{}
	input.Set("p_type", "g")

	if (input.Get("v0") != "") || (input.Get("v1") != "") {
		data, err := menuRule.List(input)
		if err != nil {
			r.Code = 5000
			r.Msg = err.Error()
		} else {
			r.Data = data
		}
	} else {
		r.Data = []string{}
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// p列表
// @Title Get User Role List
// @Description  List Roles
// @router /p [post,get]
func (this *RulesController) PList() {
	var r controllers.Returned
	input := this.Input()

	menuRule := models.MenuRule{}
	input.Set("p_type", "p")

	if input.Get("v0") != "" {
		data, err := menuRule.List(input)
		if err != nil {
			r.Code = 5000
			r.Msg = err.Error()
		} else {
			r.Data = data
		}
	} else {
		r.Data = []string{}
	}

	this.Data["json"] = r
	this.ServeJSON()
}
