package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册模型
	orm.RegisterModel( new(Rule))
}

type Rule struct {
	Id    int    `json:"id" orm:column(id);pk`
	PType string `json:"p_type" orm:"column(p_type)"`
	V0    string `json:"v0" orm:"column(v0)"`
	V1    string `json:"v1" orm:"column(v1)"`
	V2    string `json:"v2" orm:"column(v2)"`
	V3    string `json:"v3" orm:"column(v3)"`
	V4    string `json:"v4" orm:"column(v4)"`
	V5    string `json:"v5" orm:"column(v5)"`
}

// 手动设置表名
func (m *Rule) TableName() string {
	return "rules"
}
