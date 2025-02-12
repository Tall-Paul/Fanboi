// Package plugin serves as the bridge between the main application and plugins.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugin

import (
	"fmt"
)

// RoleHook takes the role contents, DB and Post and returns the text this role
// should be replaced with.
type getHook func(string) float32

// ContentsHook takes the post contents, DB and Post and returns the replacement
// contents.
type setHook func(string, float32)

type PluginManager struct {
	getHooks map[string]getHook
	setHooks map[string]setHook
}

func newPluginManager() *PluginManager {
	pm := &PluginManager{}
	pm.setHooks = make(map[string]setHook)
	pm.getHooks = make(map[string]getHook)
	return pm
}

func (pm *PluginManager) RegisterGetHook(plugin string, hook getHook) {
	pm.getHooks[plugin] = hook
}

func (pm *PluginManager) RegisterSetHook(plugin string, hook setHook) {
	pm.setHooks[plugin] = hook
}

func (pm *PluginManager) ApplyGetHook(plugin string, identifier string) (float32, error) {
	if hook, ok := pm.getHooks[plugin]; ok {
		return hook(identifier), nil
	} else {
		return -1, fmt.Errorf("no get hook for plugin '%s' found", plugin)
	}
}

func (pm *PluginManager) ApplySetHook(plugin string, identifier string, value float32) {
	if hook, ok := pm.setHooks[plugin]; ok {
		hook(identifier, value)
	}
}
