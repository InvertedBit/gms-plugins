package main

import (
	"os"
	"plugin"

	"github.com/invertedbit/gms-plugins/components"
)

type Plugin struct {
	Name        string
	Author      string
	Version     string
	Description string

	Components map[string]components.Component
}

type PluginManager struct {
	Plugins map[string]Plugin
}

func (pm *PluginManager) TryLoadPlugins(pluginDir string) error {
	if pm.Plugins == nil {
		pm.Plugins = make(map[string]Plugin)
	}

	fileList, err := os.ReadDir(pluginDir)
	if err != nil {
		return err
	}

	for _, file := range fileList {
		if file.IsDir() {
			continue
		}
		plugin, err := plugin.Open(pluginDir + file.Name())
		if err != nil {
			return err
		}
		getPluginsSymbol, err := plugin.Lookup("GetPlugins")
		if err != nil {
			return err
		}
		getPluginsFunc := getPluginsSymbol.(func() map[string]Plugin)
		plugins := getPluginsFunc()
		for name, plugin := range plugins {
			if _, exists := pm.Plugins[name]; exists {
				continue // Skip if plugin with same name already exists
			}
			pm.Plugins[name] = plugin
		}
	}
	return nil
}

func (pm *PluginManager) GetLoadedPlugins() map[string]Plugin {
	if pm.Plugins == nil {
		pm.Plugins = make(map[string]Plugin)
	}
	return pm.Plugins
}
