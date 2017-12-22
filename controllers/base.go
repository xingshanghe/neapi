package controllers

import (
	"github.com/astaxie/beego"
)


// 基础控制器
type BaseController struct {
	beego.Controller
}

//返回信息结构体
type Returned struct {
	Code int         `json:"code"` //状态码
	Msg  string      `json:"msg"`  //返回消息
	Data interface{} `json:"data"` //返回消息
}


// 认证
//func (this *BaseController) Prepare() {
//	if false {
//		this.Ctx.Redirect(302, "/")
//	}
//}