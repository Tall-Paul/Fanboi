package main

import (
	"fanboi/plugin"
	"fmt"
	"log"
)

func main() {
	pluginDir := "./plugins"
	cache := make(map[string]map[string]float32)

	pm, err := plugin.LoadPlugins(pluginDir)
	for pluginName := range pm.ListTemplatePlugins() {
		cache[pluginName] = make(map[string]float32)
	}

	if err != nil {
		log.Fatal(err)
	}

	//this bit is just to test the template stuff works
	/*fan1Temp, err := pm.ApplyGetHook("echo", "one")
	if err == nil {
		cache["template"]["fan1"] = fan1Temp
	}
	fan2Temp, err := pm.ApplyGetHook("echo", "two")
	if err == nil {
		cache["template"]["fan2"] = fan2Temp
	}
	pm.WriteTemplateHook("template", cache["template"])
	*/
	templatePlugin := pm.GetPlugin("template")
	templatePlugin.SetValue("testing", 1.00)
	templatePlugin.SetValue("testing2", 2.00)
	templatePlugin.SetValue("testing3", 3.00)

	templatePlugin2 := pm.GetPlugin("template")
	fmt.Printf("testing2 is %f", templatePlugin2.GetValue("testing2"))
}
