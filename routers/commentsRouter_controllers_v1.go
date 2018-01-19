package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:AccountsController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:AccountsController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/options`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:MenusController"],
		beego.ControllerComments{
			Method: "Tree",
			Router: `/tree`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "SetMenus",
			Router: `/setMenus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "SetUsers",
			Router: `/setUsers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RolesController"],
		beego.ControllerComments{
			Method: "Tree",
			Router: `/tree`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RulesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RulesController"],
		beego.ControllerComments{
			Method: "GList",
			Router: `/g`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RulesController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:RulesController"],
		beego.ControllerComments{
			Method: "PList",
			Router: `/p`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "EditPwd",
			Router: `/editPwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "EditSelf",
			Router: `/editSelf`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "ResetPwd",
			Router: `/resetPwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "SetRoles",
			Router: `/setRoles`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"] = append(beego.GlobalControllerRouter["github.com/xingshanghe/neapi/controllers/v1:UsersController"],
		beego.ControllerComments{
			Method: "ToggleStatus",
			Router: `/toggleStatus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
