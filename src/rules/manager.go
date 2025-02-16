package rules

import (
	"fanboi/plugin"
)

type RuleManager struct {
	rules map[int]Rule
}

type RuleInputInterface interface {
	checkInput() bool
}

type RuleOutputInterface interface {
	setOutput()
}

type RuleInput struct {
	isEmpty    bool
	plugin     plugin.PluginInterface
	identifer  string
	comparator string
	value      float32
}

type RuleOutput struct {
	plugin    plugin.PluginInterface
	identifer string
	value     float32
}

type Rule struct {
	lineNo int
	Input  RuleInput
	Output RuleOutput
}

func (rule RuleOutput) setOutput() {
	pl := rule.plugin
	pl.SetValue(rule.identifer, rule.value)
}

func (rule RuleInput) checkInput() bool {
	currentVal := rule.plugin.GetValue(rule.identifer)
	setPoint := rule.value
	switch rule.comparator {
	case "=":
		{
			if currentVal == setPoint {
				return true
			} else {
				return false
			}
		}
	case "<":
		{
			if currentVal < setPoint {
				return true
			} else {
				return false
			}
		}
	case ">":
		{
			if currentVal > setPoint {
				return true
			} else {
				return false
			}
		}
	default:
		{
			return false
		}
	}
}

func (rm *RuleManager) RunRules() {
	for _, rule := range rm.rules {
		if rule.Input.checkInput() {
			rule.Output.setOutput()
		}
	}
}
