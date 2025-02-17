// Plugin loading code.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugin

import (
	"io/fs"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

const pluginPath string = "./plugins"

func (pm *PluginManager) LoadPlugin(pluginName string) (PluginInterface, error) {
	fullpath := filepath.Join(pluginPath, pluginName+".so")
	p, err := plugin.Open(fullpath)
	if err != nil {
		return nil, err
	}

	ifunc, err := p.Lookup("InitPlugin")
	if err != nil {
		return nil, err
	}

	initFunc := ifunc.(func(*PluginManager) error)
	if err := initFunc(pm); err != nil {
		return nil, err
	}

	loadedPlugin := pm.GetPlugin(pluginName)
	loadedPlugin.StartHook()
	return loadedPlugin, nil

}

// LoadPlugins loads plugins from the directory with the given path, looking for
// all .so files in there. It creates a new PluginManager and registers the
// plugins with it.
func LoadPlugins() (*PluginManager, error) {

	entries, err := os.ReadDir(pluginPath)
	if err != nil {
		return nil, err
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	pm := NewPluginManager()

	for _, entry := range infos {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".so") {
			fullpath := filepath.Join(pluginPath, entry.Name())
			//fmt.Printf("Loaded plugin %s \n", entry.Name())

			p, err := plugin.Open(fullpath)
			if err != nil {
				return nil, err
			}

			ifunc, err := p.Lookup("InitPlugin")
			if err != nil {
				return nil, err
			}

			initFunc := ifunc.(func(*PluginManager) error)
			if err := initFunc(pm); err != nil {
				return nil, err
			}
		}
	}
	return pm, nil
}
