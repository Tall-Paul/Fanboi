package main

import (
	"fanboi/plugin"
)

func main() {

}

func InitPlugin(pm *plugin.PluginManager) error {
	pm.RegisterGetHook("unraiddrives", getDriveTemp)
	pm.RegisterSetHook("unraiddrives", setDriveTemp)
	return nil
}

func getDriveTemp(identifier string) float32 {
	return 1.0
}

func setDriveTemp(identifier string, value float32) {

}
