package models

import (
	cm "github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/xingshanghe/neapi/libs/uuid"
	"runtime"
)

// 实现 casbin.persist.CasbinMenuAdapter
type CasbinMenuAdapter struct {
	aliasName string
}

// finalizer is the destructor for CasbinMenuAdapter.
func finalizer(a *CasbinMenuAdapter) {
}

func GetCMA(alias string) *CasbinMenuAdapter {
	a := &CasbinMenuAdapter{alias}

	// Call the destructor when the object is released.
	// 手动gc
	runtime.SetFinalizer(a, finalizer)
	return a
}

// Rule对象转字符串存储
func loadPolicyLine(line MenuRule, cm cm.Model) {
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

// 字符串转Rule对象
func savePolicyLine(ptype string, rule []string) MenuRule {
	line := MenuRule{}
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
func (a *CasbinMenuAdapter) LoadPolicy(cm cm.Model) error {
	var lines []MenuRule
	err := E.Find(&lines)
	if err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, cm)
	}

	return nil
}

// 保存策略
func (a *CasbinMenuAdapter) SavePolicy(cm cm.Model) error {

	E.DropTables(&MenuRule{})
	E.CreateTables(&MenuRule{})

	var lines []MenuRule

	for ptype, ast := range cm["p"] {
		for _, MenuRule := range ast.Policy {
			line := savePolicyLine(ptype, MenuRule)
			line.Id = uuid.Rand().Raw()
			lines = append(lines, line)
		}
	}
	for ptype, ast := range cm["g"] {
		for _, MenuRule := range ast.Policy {
			line := savePolicyLine(ptype, MenuRule)
			line.Id = uuid.Rand().Raw()
			lines = append(lines, line)
		}
	}

	_, err := E.Insert(&lines)
	return err
}

// 新增策略
func (a *CasbinMenuAdapter) AddPolicy(sec string, ptype string, MenuRule []string) error {
	line := savePolicyLine(ptype, MenuRule)
	line.Id = uuid.Rand().Raw()
	_, err := E.Insert(&line)
	return err
}

// 移除策略
func (a *CasbinMenuAdapter) RemovePolicy(sec string, ptype string, MenuRule []string) error {
	line := savePolicyLine(ptype, MenuRule)
	_, err := E.Delete(&line)
	return err
}

// 根据过滤条件移除策略
func (a *CasbinMenuAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := MenuRule{}

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

	_, err := E.Delete(&line)
	return err
}
