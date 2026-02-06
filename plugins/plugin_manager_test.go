package plugins

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPluginManager_TryLoadPlugins(t *testing.T) {
	// Create a temporary directory for test plugins
	tmpDir, err := os.MkdirTemp("", "gms-plugins-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	pm := &PluginManager{}

	// Test with empty directory - should not error
	err = pm.TryLoadPlugins(tmpDir)
	if err != nil {
		t.Errorf("TryLoadPlugins with empty dir returned error: %v", err)
	}

	// Test that Plugins map is initialized
	if pm.Plugins == nil {
		t.Error("Plugins map should be initialized after TryLoadPlugins")
	}

	// Test GetLoadedPlugins returns empty map when no plugins
	plugins := pm.GetLoadedPlugins()
	if plugins == nil {
		t.Error("GetLoadedPlugins should return empty map, not nil")
	}
	if len(plugins) != 0 {
		t.Errorf("Expected 0 plugins, got %d", len(plugins))
	}
}

func TestPluginManager_GetLoadedPlugins_Uninitialized(t *testing.T) {
	pm := &PluginManager{}

	plugins := pm.GetLoadedPlugins()

	if plugins == nil {
		t.Error("GetLoadedPlugins should initialize and return empty map")
	}
	if len(plugins) != 0 {
		t.Errorf("Expected 0 plugins, got %d", len(plugins))
	}
}

func TestPluginManager_DuplicatePlugins(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "gms-plugins-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	pm := &PluginManager{}

	// Load plugins twice - should not add duplicates
	err = pm.TryLoadPlugins(tmpDir)
	if err != nil {
		t.Errorf("First TryLoadPlugins returned error: %v", err)
	}

	countBefore := len(pm.Plugins)

	err = pm.TryLoadPlugins(tmpDir)
	if err != nil {
		t.Errorf("Second TryLoadPlugins returned error: %v", err)
	}

	countAfter := len(pm.Plugins)
	if countBefore != countAfter {
		t.Errorf("Expected plugin count to remain %d after second load, got %d", countBefore, countAfter)
	}
}

func TestPluginManager_NonExistentDirectory(t *testing.T) {
	pm := &PluginManager{}

	// Try to load from non-existent directory
	err := pm.TryLoadPlugins("/non/existent/path")

	// Should not error, just return empty plugins
	if err == nil {
		t.Errorf("TryLoadPlugins on non-existent dir should error: %v", err)
	}

	if len(pm.Plugins) != 0 {
		t.Errorf("Expected 0 plugins from non-existent dir, got %d", len(pm.Plugins))
	}
}

// Helper function to create a mock .so plugin file for testing
func createMockPlugin(t *testing.T, dir, name string) string {
	t.Helper()

	// Create a minimal Go file that exports GetPlugins
	mockCode := `package main

import "github.com/invertedbit/gms-plugins/components"
import "github.com/invertedbit/gms-plugins/plugins"

func GetPlugins() map[string]plugins.Plugin {
	return map[string]plugins.Plugin{
		"test-component": {
			Name:    "Test Component",
			Author:  "Test",
			Version: "1.0.0",
			Components: map[string]components.Component{
				"test": {},
			},
		},
	}
}
`

	// Write the mock plugin file
	pluginFile := filepath.Join(dir, name+".go")
	if err := os.WriteFile(pluginFile, []byte(mockCode), 0644); err != nil {
		t.Fatalf("Failed to write mock plugin: %v", err)
	}

	return pluginFile
}
