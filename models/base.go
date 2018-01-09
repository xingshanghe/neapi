package models

import (
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"

	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/xingshanghe/neapi/libs"
)

var E *xorm.Engine
var Cma *CasbinMenuAdapter
var Cme *casbin.Enforcer

//返回信息结构体
type Paged struct {
	Total    int64 `json:"total"`
	PageSize int   `json:"page_size"`
	PageNo   int   `json:"page_no"`
}

func init() {
	initXorm("default")

	appConf, _ := GetAppConf()
	Cma = GetCMA("default")
	Cme = casbin.NewEnforcer("conf/rbac-menu.conf", Cma, appConf.String("runmode") == "dev")
	Cme.LoadPolicy()
	//casbin.NewEnforcer("examples/rbac_model.conf", "")
}

func GetPassword(password string) string {
	appConf, _ := GetAppConf()
	return libs.MD5(libs.MD5(password) + appConf.String("salt"))
}

func initXorm(alias string) {
	appConf, _ := GetAppConf()
	max, min := getMysqlConn()
	mysqlResource := getMysqlResource()

	engine, _ := xorm.NewEngine("mysql", mysqlResource)
	engine.SetMaxIdleConns(min)
	engine.SetMaxOpenConns(max)

	engine.ShowSQL(appConf.String("runmode") == "dev")
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	E = engine
}

func initOrmAndXorm(alias string) {
	appConf, _ := GetAppConf()
	max, min := getMysqlConn()
	mysqlResource := getMysqlResource()

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(alias, "mysql", mysqlResource, max, min)
	orm.Debug = appConf.String("runmode") == "dev"

	engine, _ := xorm.NewEngine("mysql", mysqlResource)
	engine.SetMaxIdleConns(min)
	engine.SetMaxOpenConns(max)

	engine.ShowSQL(appConf.String("runmode") == "dev")
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	E = engine
}

func GetAppConf() (config.Configer, error) {
	return config.NewConfig("ini", "conf/app.conf")
}
func getMysqlResource() string {
	conf, _ := config.NewConfig("ini", "conf/store.conf")
	return conf.String("mysql::source")
}
func getMysqlConn() (int, int) {
	conf, _ := config.NewConfig("ini", "conf/store.conf")
	max, _ := conf.Int("mysql::connMax")
	min, _ := conf.Int("mysql::connMin")
	return max, min
}
