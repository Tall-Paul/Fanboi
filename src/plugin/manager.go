// Package plugin serves as the bridge between the main application and plugins.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugin

import (
	"fmt"
)

type PluginTemplateData struct {
	Identifier string
	Value      float32
}

type getHook func(string) float32

// the 'set' hook takes an identifier and a value (eg: fan1, 50)
type setHook func(string, float32)

type isTemplateHook func() bool

type writeTemplateHook func(data map[int]PluginTemplateData)

type PluginManager struct {
	getHooks           map[string]getHook
	setHooks           map[string]setHook
	isTemplateHooks    map[string]isTemplateHook
	writeTemplateHooks map[string]writeTemplateHook
}

func newPluginManager() *PluginManager {
	pm := &PluginManager{}
	pm.setHooks = make(map[string]setHook)
	pm.getHooks = make(map[string]getHook)
	pm.writeTemplateHooks = make(map[string]writeTemplateHook)
	pm.isTemplateHooks = make(map[string]isTemplateHook)
	return pm
}

func (pm *PluginManager) RegisterGetHook(plugin string, hook getHook) {
	pm.getHooks[plugin] = hook
}

func (pm *PluginManager) RegisterSetHook(plugin string, hook setHook) {
	pm.setHooks[plugin] = hook
}

func (pm *PluginManager) RegisterWriteTemplateHook(plugin string, hook writeTemplateHook) {
	pm.writeTemplateHooks[plugin] = hook
}

func (pm *PluginManager) RegisterisTemplateHook(plugin string, hook isTemplateHook) {
	pm.isTemplateHooks[plugin] = hook
}

func (pm *PluginManager) CheckIsTemplateHook(plugin string) bool {
	if hook, ok := pm.isTemplateHooks[plugin]; ok {
		return hook()
	} else {
		fmt.Printf("no get hook for plugin %s' found", plugin)
		return false
	}
}

func (pm *PluginManager) WriteTemplateHook(plugin string, data map[int]PluginTemplateData) {
	if hook, ok := pm.writeTemplateHooks[plugin]; ok {
		hook(data)
	} else {
		fmt.Printf("no get hook for plugin %s' found", plugin)
	}
}

func (pm *PluginManager) ApplyGetHook(plugin string, identifier string) (float32, error) {
	if hook, ok := pm.getHooks[plugin]; ok {
		return hook(identifier), nil
	} else {
		return -1, fmt.Errorf("no get hook for plugin %s' found", plugin)
	}
}

func (pm *PluginManager) ApplySetHook(plugin string, identifier string, value float32) {
	if hook, ok := pm.setHooks[plugin]; ok {
		hook(identifier, value)
	}
}
