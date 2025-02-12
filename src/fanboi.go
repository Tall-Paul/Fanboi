package main

import (
	"fanboi/plugin"
	"fmt"
	"log"
)

func main() {
	pluginDir := "./plugins"

	pm, err := plugin.LoadPlugins(*&pluginDir)
	if err != nil {
		log.Fatal(err)
	}
	temperature, err := pm.ApplyGetHook("unraid_drives", "drive1")
	if err == nil {
		fmt.Printf("Plugin returned %f", temperature)
	} else {
		log.Fatal(err)
	}

}
