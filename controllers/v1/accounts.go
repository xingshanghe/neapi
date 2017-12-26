package v1

import (
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
	//captcha, _ := dataBody.Get("captcha").String()

	var data struct {
		Token   string      `json:"token"`
		Account models.User `json:"account"`
	}
	account := models.Account{
		Username: username,
		Password: password,
	}
	var hasThisAccount bool
	hasThisAccount, err = models.E.Get(&account)
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
			var token string
			token, err = libs.CreateJwt(user)

			if err == nil {
				data.Token = token
				data.Account = user
				r.Data = data
			} else {
				r.Data = new(struct {
					Token string `json:"token"`
				})
				r.Msg = err.Error()
			}
		} else {
			r.Data = new(struct {
				Token string `json:"token"`
			})
			r.Code = 5000
			r.Msg = "账号密码验证错误."
		}
	} else {
		r.Data = new(struct {
			Token string `json:"token"`
		})
		r.Msg = err.Error()
	}
	this.Data["json"] = r
	this.ServeJSON()
}
