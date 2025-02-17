package rules

import (
	"fanboi/plugin"
	"fmt"
)

type RuleManager struct {
	pm    *plugin.PluginManager
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
	if rule.isEmpty {
		return true
	}
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
	for _, pl := range rm.pm.GetPlugins() {
		pl.StartHook()
	}
	for i := 1; i <= len(rm.rules); i++ {
		rule := rm.rules[i]
		if rule.Input.checkInput() {
			if !rule.Input.isEmpty {
				fmt.Printf("[%v] Updating %s to %v", rule.lineNo, rule.Output.identifer, rule.Output.value)
				fmt.Println()
			} else {
				fmt.Printf("[%v] defaulting %s to %v", rule.lineNo, rule.Output.identifer, rule.Output.value)
				fmt.Println()
			}
			rule.Output.setOutput()
		} else {
		}
	}
	for _, pl := range rm.pm.GetPlugins() {
		pl.EndHook()
	}
}

func (rm *RuleManager) DumpRules() {
	fmt.Printf("%+v\n", rm.rules)
}
