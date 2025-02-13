package main

import (
	"fanboi/plugin"
	"fmt"
	"strconv"
	"strings"
)

func InitPlugin(pm *plugin.PluginManager) error {

	pm.RegisterisTemplateHook("template", returnIsTemplate)
	pm.RegisterWriteTemplateHook("template", writeTemplate)
	return nil
}

func returnIsTemplate() bool {
	return true
}

func writeTemplate(data map[int]plugin.PluginTemplateData) {
	template := "fan1: {fan1}, fan2: {fan2}"
	for _, value := range data {
		textValue := strconv.FormatFloat(float64(value.Value), 'f', -1, 32)
		template = strings.Replace(template, "{"+value.Identifier+"}", textValue, -1)
	}
	fmt.Println(template)
}
