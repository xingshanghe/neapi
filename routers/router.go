// @APIVersion 1.0.0
// @Title dop cloud API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact xingshanghe@gmail.com
// @TermsOfServiceUrl http://sgrcloud.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/xingshanghe/neapi/controllers/v1"
)

func init() {

	dopns := beego.NewNamespace("/v1",
		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&v1.AccountsController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&v1.UsersController{},
			),
		),
		beego.NSNamespace("/roles",
			beego.NSInclude(
				&v1.RolesController{},
			),
		),
		beego.NSNamespace("/rules",
			beego.NSInclude(
				&v1.RulesController{},
			),
		),
	)

	beego.AddNamespace(dopns)
}
