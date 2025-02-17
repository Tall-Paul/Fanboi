package main

import (
	"fanboi/plugin"
	"fanboi/rules"
)

func main() {

	pm := plugin.NewPluginManager()

	rulesFile := "./rules.fnb"
	ok, rm := rules.LoadRules(rulesFile, pm)
	if ok {
		rm.RunRules()
	}
	for _, pl := range pm.GetPlugins() {
		pl.EndHook()
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
