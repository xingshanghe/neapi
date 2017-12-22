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
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}