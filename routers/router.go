// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/xingshanghe/neapi/controllers"
	"github.com/xingshanghe/neapi/controllers/v1"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/test",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)

	nens := beego.NewNamespace("/v1",
		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&v1.AccountsController{},
			),
		),
	)

	beego.AddNamespace(ns, nens)
}
