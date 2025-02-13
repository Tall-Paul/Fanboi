package main

import (
	"fanboi/plugin"
	"fmt"
	"strconv"
	"strings"
)

func InitPlugin(pm *plugin.PluginManager) error {

	pm.RegisterWriteTemplateHook("template", writeTemplate)
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
