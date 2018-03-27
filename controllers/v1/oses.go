package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type OsesController struct {
	controllers.BaseController
}

// 列表
// @Title Get Oses List
// @Description  List Oses
// @router / [get]
func (this *OsesController) Get() {
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
// @router / [post]
func (this *OsesController) Post() {
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
// @router /:id [put]
func (this *OsesController) Put() {
	var r controllers.Returned

	input := this.Input()

	id := this.Ctx.Input.Param(":id")
	os := models.Os{
		Id: id,
	}
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
// @router /:id [delete]
func (this *OsesController) Delete() {
	var r controllers.Returned

	input := this.Input()

	id := this.Ctx.Input.Param(":id")
	os := models.Os{
		Id: id,
	}
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
// @router /options [get]
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
