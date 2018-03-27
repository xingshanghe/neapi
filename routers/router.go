// @APIVersion 1.0.0
// @Title CMOP Console API
// @Description CMOP Console API
// @Contact xingshanghe@gmail.com
// @TermsOfServiceUrl http://localhost:8888/
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
		beego.NSNamespace("/menus",
			beego.NSInclude(
				&v1.MenusController{},
			),
		),
		beego.NSNamespace("/regions",
			beego.NSInclude(
				&v1.RegionsController{},
			),
		),
		beego.NSNamespace("/nodes",
			beego.NSInclude(
				&v1.NodesController{},
			),
		),
		beego.NSNamespace("/oses",
			beego.NSInclude(
				&v1.OsesController{},
			),
		),
		beego.NSNamespace("/clusters",
			beego.NSInclude(
				&v1.ClustersController{},
			),
		),
	)

	beego.AddNamespace(dopns)
}
