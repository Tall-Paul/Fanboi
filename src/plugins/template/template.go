package main

import (
	"fanboi/plugin"
	"fmt"
	"strconv"
	"strings"
)

type templatePlugin struct {
	template string
	values   map[string]float32
}

func (tem templatePlugin) SetValue(identifier string, value float32) {
	tem.values[identifier] = value
}

func (tem templatePlugin) GetValue(identifier string) float32 {
	return tem.values[identifier]
}

func (tem templatePlugin) StartHook() {

}

func (tem templatePlugin) EndHook() {

}

func InitPlugin(pm *plugin.PluginManager) error {
	values := make(map[string]float32)
	this := templatePlugin{"", values}
	pm.RegisterWriteTemplateHook("template", writeTemplate)
	pm.RegisterPlugin("template", this)
	return nil
}

func writeTemplate(data map[string]float32) {
	template := "fan1: {fan1}, fan2: {fan2}"
	for key, value := range data {
		textValue := strconv.FormatFloat(float64(value), 'f', -1, 32)
		template = strings.Replace(template, "{"+key+"}", textValue, -1)
	}
	fmt.Println(template)
}
