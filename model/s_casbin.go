package model

import (
	"ginserver/module/db"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
)

type Casbin struct {
	*db.Model `xorm:"-"`
	PType     string `xorm:"varchar(20) not null default '''' index"`
	V0        string `xorm:"varchar(50) not null default '''' index"`
	V1        string `xorm:"varchar(100) not null default '''' index"`
	V2        string `xorm:"varchar(100) not null default '''' index"`
	V3        string `xorm:"varchar(100) not null default '''' index"`
	V4        string `xorm:"varchar(100) not null default '''' index"`
	V5        string `xorm:"varchar(100) not null default '''' index"`
}

func (c *Casbin) TableName() string {
	return "s_casbin"
}

func (c *Casbin) LoadPolicy(model model.Model) error {
	var rules []*Casbin
	if err := c.Select(c, &rules); err != nil {
		return err
	}
	for _, rule := range rules {
		loadPolicyLine(rule, model)
	}
	return nil
}

// SavePolicy saves policy to database.
func (c *Casbin) SavePolicy(model model.Model) error {
	var rules []*Casbin

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			rules = append(rules, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			rules = append(rules, line)
		}
	}

	return c.Insert(&rules)
}

// AddPolicy adds a policy rule to the storage.
func (c *Casbin) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	return c.InsertOne(line)
}

// RemovePolicy removes a policy rule from the storage.
func (c *Casbin) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	return c.Delete(line)
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (c *Casbin) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	rule := &Casbin{PType: ptype}

	idx := fieldIndex + len(fieldValues)
	switch {
	case fieldIndex <= 0 && idx > 0:
		rule.V0 = fieldValues[0-fieldIndex]
		fallthrough
	case fieldIndex <= 1 && idx > 1:
		rule.V1 = fieldValues[1-fieldIndex]
		fallthrough
	case fieldIndex <= 2 && idx > 2:
		rule.V2 = fieldValues[2-fieldIndex]
		fallthrough
	case fieldIndex <= 3 && idx > 3:
		rule.V3 = fieldValues[3-fieldIndex]
		fallthrough
	case fieldIndex <= 4 && idx > 4:
		rule.V4 = fieldValues[4-fieldIndex]
		fallthrough
	case fieldIndex <= 5 && idx > 5:
		rule.V5 = fieldValues[5-fieldIndex]
	}

	return c.Delete(rule)
}

const prefixLine = ", "

func loadPolicyLine(rule *Casbin, model model.Model) {
	line := rule.PType
	switch {
	case len(rule.V0) > 0:
		line += prefixLine + rule.V0
		fallthrough
	case len(rule.V1) > 0:
		line += prefixLine + rule.V1
		fallthrough
	case len(rule.V2) > 0:
		line += prefixLine + rule.V2
		fallthrough
	case len(rule.V3) > 0:
		line += prefixLine + rule.V3
		fallthrough
	case len(rule.V4) > 0:
		line += prefixLine + rule.V4
		fallthrough
	case len(rule.V5) > 0:
		line += prefixLine + rule.V5
	}
	persist.LoadPolicyLine(line, model)
}

func savePolicyLine(ptype string, rule []string) *Casbin {
	line := &Casbin{PType: ptype}

	l := len(rule)
	switch {
	case l > 0:
		line.V0 = rule[0]
		fallthrough
	case l > 1:
		line.V1 = rule[1]
		fallthrough
	case l > 2:
		line.V2 = rule[2]
		fallthrough
	case l > 3:
		line.V3 = rule[3]
		fallthrough
	case l > 4:
		line.V4 = rule[4]
		fallthrough
	case l > 5:
		line.V5 = rule[5]
	}

	return line
}
