package plugins

import (
	"context"
	"os"
	"plugin"

	"github.com/invertedbit/gms-plugins/components"
)

// HookHandler is a function that handles a hook event
type HookHandler func(ctx context.Context, args map[string]interface{}) error

// Plugin represents a loaded plugin
type Plugin struct {
	Name        string
	Author      string
	Version     string
	Description string
	Permissions []string

	Components map[string]components.Component
	Hooks      map[string]HookHandler
}

// PluginManager manages plugin loading and hook registration
type PluginManager struct {
	Plugins    map[string]Plugin
	HookRegistry
}

// HookRegistry manages hook registration and execution
type HookRegistry struct {
	hooks map[string][]HookRegistration
}

// HookRegistration represents a registered hook
type HookRegistration struct {
	PluginID string
	Priority int
	Handler  HookHandler
}

// NewPluginManager creates a new PluginManager with an empty HookRegistry
func NewPluginManager() *PluginManager {
	return &PluginManager{
		Plugins: make(map[string]Plugin),
		HookRegistry: HookRegistry{
			hooks: make(map[string][]HookRegistration),
		},
	}
}

// NewHookRegistry creates a new HookRegistry
func NewHookRegistry() *HookRegistry {
	return &HookRegistry{
		hooks: make(map[string][]HookRegistration),
	}
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

// RegisterHook registers a hook handler for a specific hook name
func (hr *HookRegistry) RegisterHook(pluginID, hookName string, priority int, handler HookHandler) {
	registration := HookRegistration{
		PluginID: pluginID,
		Priority: priority,
		Handler:  handler,
	}
	hr.hooks[hookName] = append(hr.hooks[hookName], registration)

	// Sort by priority (higher priority first)
	for i := len(hr.hooks[hookName]) - 1; i > 0; i-- {
		if hr.hooks[hookName][i].Priority > hr.hooks[hookName][i-1].Priority {
			hr.hooks[hookName][i], hr.hooks[hookName][i-1] = hr.hooks[hookName][i-1], hr.hooks[hookName][i]
		}
	}
}

// UnregisterHooks removes all hooks registered by a specific plugin
func (hr *HookRegistry) UnregisterHooks(pluginID string) {
	for hookName, registrations := range hr.hooks {
		var filtered []HookRegistration
		for _, reg := range registrations {
			if reg.PluginID != pluginID {
				filtered = append(filtered, reg)
			}
		}
		hr.hooks[hookName] = filtered
	}
}

// GetHooks returns all registered hooks for a specific hook name
func (hr *HookRegistry) GetHooks(hookName string) []HookRegistration {
	return hr.hooks[hookName]
}

// ExecuteHooks executes all handlers for a specific hook name
func (hr *HookRegistry) ExecuteHooks(ctx context.Context, hookName string, args map[string]interface{}) []error {
	var errors []error
	registrations := hr.hooks[hookName]
	for _, reg := range registrations {
		if err := reg.Handler(ctx, args); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// GetAllHooks returns all registered hooks
func (hr *HookRegistry) GetAllHooks() map[string][]HookRegistration {
	return hr.hooks
}
