package models

import (
	"runtime"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"

	cm "github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var E *xorm.Engine

//返回信息结构体
type Paged struct {
	Total    int64 `json:"total"`
	PageSize int   `json:"page_size"`
	PageNo   int   `json:"page_no"`
}

func init() {
	initOrmAndXorm("default")
	//casbin.NewEnforcer("examples/rbac_model.conf", "")
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

// 实现 casbin.persist.CasbinAdapter
type CasbinAdapter struct {
	aliasName string
	o         orm.Ormer
}

// finalizer is the destructor for CasbinAdapter.
func finalizer(a *CasbinAdapter) {
}

func GetA(alias string) *CasbinAdapter {
	a := &CasbinAdapter{alias, orm.NewOrm()}

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)
	return a
}

func loadPolicyLine(line Rule, cm cm.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}

	persist.LoadPolicyLine(lineText, cm)
}

func savePolicyLine(ptype string, rule []string) Rule {
	line := Rule{}
	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return line
}

// 加载所有策略
func (a *CasbinAdapter) LoadPolicy(cm cm.Model) error {
	var lines []Rule
	_, err := a.o.QueryTable("rule").All(&lines)
	if err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, cm)
	}
	return nil
}

// 保存策略
func (a *CasbinAdapter) SavePolicy(cm cm.Model) error {
	var lines []Rule

	for ptype, ast := range cm["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range cm["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	_, err := a.o.InsertMulti(len(lines), lines)
	return err
}

// 新增策略
func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.o.Insert(&line)
	return err
}

// 移除策略
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.o.Delete(&line)
	return err
}

// 根据过滤条件移除策略
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := Rule{}

	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}

	_, err := a.o.Delete(&line)
	return err
}
