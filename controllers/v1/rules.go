package v1
import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/models"
)

type RulesController struct {
	controllers.BaseController
}


// 列表
// @Title Get User Role List
// @Description  List Roles
// @router / [post,get]
func (this *RulesController) List() {
	var r controllers.Returned
	input := this.Input()

	rule := models.Rule{}
	data, err := rule.List(input)
	if err != nil {
		r.Code = 5000
		r.Msg = err.Error()
	} else {
		r.Data = data
	}

	this.Data["json"] = r
	this.ServeJSON()
}