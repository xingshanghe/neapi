package v1

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type RegionsController struct {
	controllers.BaseController
}

// 列表
// @Title Get Reions List
// @Description  List Roles
// @router / [get]
func (this *RegionsController) Get() {
	var r controllers.Returned
	input := this.Input()

	region := models.Region{}
	data, err := region.Page(input)
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
func (this *RegionsController) Post() {
	var r controllers.Returned

	input := this.Input()

	region := models.Region{}
	err := region.Add(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = region
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 编辑
// @Title Edit a Region
// @Description  Edit Region
// @router /:id [put]
func (this *RegionsController) Put() {
	var r controllers.Returned

	input := this.Input()

	id := this.Ctx.Input.Param(":id")
	region := models.Region{
		Id:id,
	}
	err := region.Edit(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = region
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 删除
// @Title Delete a Region
// @Description  Delete Region
// @router /:id [delete]
func (this *RegionsController) Delete() {
	var r controllers.Returned

	input := this.Input()

	id := this.Ctx.Input.Param(":id")
	region := models.Region{
		Id:id,
	}
	err := region.Delete(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = region
	}

	this.Data["json"] = r
	this.ServeJSON()
}

// 树状结构菜单列表
// @Title Get Region Options
// @Description  Get Region Options
// @router /options [get]
func (this *RegionsController) Options() {
	var r controllers.Returned
	data, err := models.RegionOptionList()
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}