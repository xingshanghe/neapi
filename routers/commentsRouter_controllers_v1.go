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

}
