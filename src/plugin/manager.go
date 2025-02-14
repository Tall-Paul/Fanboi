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

type PluginInterface interface {
	GetValue(indentifier string) float32
	SetValue(identifier string, value float32)
	StartHook()
	EndHook()
}

type getHook func(string) float32

// the 'set' hook takes an identifier and a value (eg: fan1, 50)
type setHook func(string, float32)

type writeTemplateHook func(data map[string]float32)

type PluginManager struct {
	getHooks           map[string]getHook
	setHooks           map[string]setHook
	writeTemplateHooks map[string]writeTemplateHook
	plugins            map[string]PluginInterface
}

func newPluginManager() *PluginManager {
	pm := &PluginManager{}
	pm.setHooks = make(map[string]setHook)
	pm.getHooks = make(map[string]getHook)
	pm.writeTemplateHooks = make(map[string]writeTemplateHook)
	pm.plugins = make(map[string]PluginInterface)
	return pm
}

func (pm *PluginManager) RegisterPlugin(plugin string, pluginInterface PluginInterface) {
	pm.plugins[plugin] = pluginInterface
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

func (pm *PluginManager) WriteTemplateHook(plugin string, data map[string]float32) {
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

func (pm *PluginManager) ListTemplatePlugins() map[string]bool {
	out := make(map[string]bool)
	for pluginName := range pm.writeTemplateHooks {
		out[pluginName] = true
	}
	return out
}

func (pm *PluginManager) GetPlugin(pluginName string) PluginInterface {
	return pm.plugins[pluginName]
}
