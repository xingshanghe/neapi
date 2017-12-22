package libs

import (
"github.com/astaxie/beego/config"
"github.com/astaxie/beego/logs"
)

var (
	Logger *logs.BeeLogger
)

func init() {
	go initLogger()
}

func initLogger() {

	appConf, _ := config.NewConfig("ini", "conf/app.conf")

	Logger = logs.NewLogger(LogLens)
	Logger.SetLogger("file", `{"filename":"tmp/logs/`+appConf.String("appname")+`.log"}`)
}
