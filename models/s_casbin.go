package models

import (
	"ginserver/modules/db"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
)

type SCasbin struct {
	*db.Model `xorm:"-"`
	Id        int    `xorm:"int(11) pk autoincr"`
	Ptype     string `xorm:"varchar(255) index not null default ''"`
	V0        string `xorm:"varchar(255) index not null default ''"`
	V1        string `xorm:"varchar(255) index not null default ''"`
	V2        string `xorm:"varchar(255) index not null default ''"`
	V3        string `xorm:"varchar(255) index not null default ''"`
	V4        string `xorm:"varchar(255) index not null default ''"`
	V5        string `xorm:"varchar(255) index not null default ''"`
}

func (p *SCasbin) TableName() string {
	return "s_casbin"
}

func (p *SCasbin) LoadPolicy(model model.Model) error {
	var rules []*SCasbin
	if err := p.Select(p, &rules); err != nil {
		return err
	}
	for _, line := range rules {
		loadPolicyLine(line, model)
	}
	return nil
}

// SavePolicy saves policy to database.
func (p *SCasbin) SavePolicy(model model.Model) error {
	var rules []*SCasbin

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

	return p.Insert(&rules)
}

// AddPolicy adds a policy rule to the storage.
func (p *SCasbin) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	return p.InsertOne(line)
}

// RemovePolicy removes a policy rule from the storage.
func (p *SCasbin) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	return p.Delete(line)
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (p *SCasbin) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	rule := &SCasbin{Ptype: ptype}

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

	return p.Delete(rule)
}

const prefixLine = ", "

func loadPolicyLine(rule *SCasbin, model model.Model) {
	line := rule.Ptype
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

func savePolicyLine(ptype string, rule []string) *SCasbin {
	line := &SCasbin{Ptype: ptype}

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
