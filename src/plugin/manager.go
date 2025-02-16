// Package plugin serves as the bridge between the main application and plugins.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugin

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

type PluginManager struct {
	plugins map[string]PluginInterface
}

func newPluginManager() *PluginManager {
	pm := &PluginManager{}
	pm.plugins = make(map[string]PluginInterface)
	return pm
}

func (pm *PluginManager) RegisterPlugin(plugin string, pluginInterface PluginInterface) {
	pm.plugins[plugin] = pluginInterface
}

func (pm *PluginManager) GetPlugin(pluginName string) PluginInterface {
	pl, ok := pm.plugins[pluginName]
	if ok {
		return pl
	} else {
		return nil
	}
}

func (pm *PluginManager) GetPlugins() map[string]PluginInterface {
	return pm.plugins
}
