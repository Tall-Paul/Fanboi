package main

import (
	"fanboi/plugin"
)

type echoPlugin struct {
}

func (pl echoPlugin) SetValue(identifier string, value float32) {
	//noop
}

func (pl echoPlugin) GetValue(identifier string) float32 {
	switch identifier {
	case "one":
		return 1.00
	case "two":
		return 2.00
	default:
		return 100.00
	}
}

func (pl echoPlugin) StartHook() {

}

func (pl echoPlugin) EndHook() {

}

func InitPlugin(pm *plugin.PluginManager) error {
	pl := echoPlugin{}
	pm.RegisterPlugin("echo", pl)
	return nil
}
