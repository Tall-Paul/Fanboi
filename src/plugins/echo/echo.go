package main

import (
	"fanboi/plugin"
	"fmt"
)

func InitPlugin(pm *plugin.PluginManager) error {
	pm.RegisterSetHook("echo", outputText)
	pm.RegisterGetHook("echo", echoValue)
	return nil
}

func echoValue(identifier string) float32 {
	switch identifier {
	case "one":
		return 1.00
	case "two":
		return 2.00
	default:
		return 100.00
	}
}

func outputText(identifier string, value float32) {
	fmt.Printf("Identifier %s set to %f", identifier, value)
}
