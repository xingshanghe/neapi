package v1

import (
	"net/url"
	"strings"

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
// @router /login [post]
func (this *AccountsController) Login() {
	var r controllers.Returned

	dataBody, err := simplejson.NewJson(this.Ctx.Input.RequestBody)

	username, _ := dataBody.Get("username").String()
	password, _ := dataBody.Get("password").String()
	//captcha, _ := dataBody.Get("captcha").String()

	var data struct {
		Token   string         `json:"token"`
		Account models.User    `json:"account"`
		Roles   []string       `json:"roles"`
		Menus   []*models.Menu `json:"menus"`
	}
	account := models.Account{
		Username: username,
		Password: models.GetPassword(password),
	}
	var hasThisAccount bool
	hasThisAccount, err = models.E.Get(&account)
	if err == nil {
		if hasThisAccount {

			if account.Status == 0 {
				detail := models.Detail{
					AccountId: account.Id,
				}
				models.E.Get(&detail)
				user := models.User{
					account,
					detail,
				}

				roleIds := []string{}

				menuRule := models.MenuRule{}
				p := url.Values{}
				p.Set("p_type", "g")
				p.Set("v0", account.Id)
				menuRules, _ := menuRule.List(p)
				for _, mr := range menuRules {
					roleIds = append(roleIds, mr.V1)
				}

				var token string
				token, err = libs.CreateJwt(user, roleIds)

				menuIds := []string{}
				p1 := url.Values{}
				p1.Set("p_type", "p")
				p1.Set("v0", strings.Join(roleIds, ","))
				menuRules2, _ := menuRule.List(p1)

				for _, mr := range menuRules2 {
					menuIds = append(menuIds, mr.V1)
				}

				// 获取左边侧边栏菜单
				p2 := url.Values{}
				p2.Set("is_side", "1")
				menuTree, err := models.GetMenusTree("", menuIds, true, p2)

				if err == nil {
					data.Roles = roleIds
					data.Token = token
					data.Account = user
					data.Menus = menuTree
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
				r.Code = 5001
				r.Msg = "该帐号已被锁定禁止登录."
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
