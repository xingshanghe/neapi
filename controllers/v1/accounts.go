package v1

import (
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"

	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/libs"
	"github.com/xingshanghe/neapi/models"
)

type AccountsController struct {
	controllers.BaseController
}

// 登录
// @Title Login
// @Description  Login Users
// @router /login [post,get]
func (this *AccountsController) Login() {
	var r controllers.Returned

	dataBody, err := simplejson.NewJson(this.Ctx.Input.RequestBody)

	username, _ := dataBody.Get("username").String()
	password, _ := dataBody.Get("password").String()
	captcha, _ := dataBody.Get("captcha").String()

	logs.Info(username, password, captcha)

	var data struct {
		Token   string         `json:"token"`
		Account models.Account `json:"account"`
	}
	//TODO 验证用户，密码
	var token string
	account := models.Account{
		Username: username,
		Password: password,
		Profile:  models.Detail{"邢尚合", "男", 32, "四川省成都市", username + "@gmail.com"},
	}
	if err == nil {
		token, err = libs.CreateJwt(account)

		if err == nil {
			data.Token = token
			data.Account = account

		} else {
			r.Msg = err.Error()
		}

	} else {
		r.Msg = err.Error()
	}

	r.Data = data

	this.Data["json"] = r
	this.ServeJSON()
}
