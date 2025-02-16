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
	out := tem.template
	for key, value := range tem.values {
		textValue := strconv.FormatFloat(float64(value), 'f', -1, 32)
		out = strings.Replace(out, "{"+key+"}", textValue, -1)
	}
	fmt.Println(out)
}

func InitPlugin(pm *plugin.PluginManager) error {
	values := make(map[string]float32)
	this := templatePlugin{"fan1: {fan1_speed}, fan2: {fan2_speed}", values}
	pm.RegisterPlugin("template", this)
	return nil
}
