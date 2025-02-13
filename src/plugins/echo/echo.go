package main

import (
	"fanboi/plugin"
	"fmt"
	"strconv"
)

func main() {

}

func InitPlugin(pm *plugin.PluginManager) error {
	pm.RegisterSetHook("echo", outputText)
	pm.RegisterGetHook("echo", echoValue)
	return nil
}

func echoValue(identifier string) float32 {
	returnval, err := strconv.ParseFloat(identifier, 32)
	if err == nil {
		return float32(returnval)
	} else {
		return -1
	}
}

func outputText(identifier string, value float32) {
	fmt.Printf("Identifier %s set to %f", identifier, value)
}
