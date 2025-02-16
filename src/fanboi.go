package main

import (
	"fanboi/plugin"
	"fanboi/rules"
	"fmt"
	"log"
)

func main() {
	pluginDir := "./plugins"
	pm, err := plugin.LoadPlugins(pluginDir)
	if err != nil {
		log.Fatal(err)
	}
	rulesFile := "./rules.fnb"
	ok, rm := rules.LoadRules(rulesFile, pm)
	if ok {
		for _, pl := range pm.GetPlugins() {
			pl.StartHook()
		}

		rm.RunRules()

		for _, pl := range pm.GetPlugins() {
			pl.EndHook()
		}
	} else {
		fmt.Println("Error loading rules")
	}

	/*
		fmt.Println()

		//check unraiddrive plugin works
		ud := pm.GetPlugin("unraiddrives")
		fmt.Printf("parity temp is %f", ud.GetValue("parity"))
		fmt.Println()

		//check template plugin works
		templatePlugin := pm.GetPlugin("template")
		templatePlugin.SetValue("fan1", 5.00)
		templatePlugin.SetValue("fan2", 10.00)
		templatePlugin.SetValue("fan3", 15.00)
	*/

}
