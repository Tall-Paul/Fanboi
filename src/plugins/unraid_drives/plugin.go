package main

import (
	"fanboi/plugin"
)

func InitPlugin(pm *plugin.PluginManager) error {
	pm.RegisterGetHook("unraid_drives", getDriveTemp)
	pm.RegisterSetHook("unraid_drives", setDriveTemp)
	return nil
}

func getDriveTemp(identifier string) float32 {
	return 1.0
}

func setDriveTemp(identifier string, value float32) {

}
