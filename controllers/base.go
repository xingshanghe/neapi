package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"os"
)

const SYSTEM_LOG_FILE = "tmp/logs/neapi.log"

var NeLogger *logs.BeeLogger

func init() {
	go initNeLogger()
}

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

// 初始化G2logger
func initNeLogger() {
	NeLogger = logs.NewLogger(10000)
	//os.Create(SYSTEM_LOG_FILE)
	NeLogger.SetLogger("file", `{"filename":"`+SYSTEM_LOG_FILE+`"}`)
}

// 认证
//func (this *BaseController) Prepare() {
//	if false {
//		this.Ctx.Redirect(302, "/")
//	}
//}