package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// PluginManifest represents the plugin manifest.json file
type PluginManifest struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Author       string            `json:"author"`
	Description  string            `json:"description"`
	GMSVersion   string            `json:"gmsVersion"`
	Permissions  []string          `json:"permissions"`
	Hooks        []HookDefinition  `json:"hooks"`
	Components   []ComponentDef    `json:"components"`
	Assets       AssetDefinition   `json:"assets"`
}

// HookDefinition defines a hook that the plugin provides
type HookDefinition struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// ComponentDef defines a component provided by the plugin
type ComponentDef struct {
	Slug        string             `json:"slug"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Properties  []PropertyDef      `json:"properties"`
}

// PropertyDef defines a component property
type PropertyDef struct {
	Slug     string `json:"slug"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

// AssetDefinition defines static assets included in the plugin
type AssetDefinition struct {
	CSS []string `json:"css"`
	JS  []string `json:"js"`
}

// Load loads a plugin manifest from a file path
func Load(path string) (*PluginManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest PluginManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

// LoadFromDirectory loads a plugin manifest from a directory
func LoadFromDirectory(dir string) (*PluginManifest, error) {
	manifestPath := filepath.Join(dir, "manifest.json")
	return Load(manifestPath)
}

// Validate validates the manifest structure
func (m *PluginManifest) Validate() error {
	if m.Name == "" {
		return ErrMissingName
	}
	if m.Version == "" {
		return ErrMissingVersion
	}
	return nil
}

// HasHook checks if the plugin provides a specific hook
func (m *PluginManifest) HasHook(hookName string) bool {
	for _, h := range m.Hooks {
		if h.Name == hookName {
			return true
		}
	}
	return false
}

// HasComponent checks if the plugin provides a specific component
func (m *PluginManifest) HasComponent(slug string) bool {
	for _, c := range m.Components {
		if c.Slug == slug {
			return true
		}
	}
	return false
}

// HasPermission checks if the plugin has a specific permission
func (m *PluginManifest) HasPermission(perm string) bool {
	for _, p := range m.Permissions {
		if p == perm {
			return true
		}
	}
	return false
}

// Errors
var (
	ErrMissingName    = &ValidationError{Field: "name", Message: "plugin name is required"}
	ErrMissingVersion = &ValidationError{Field: "version", Message: "plugin version is required"}
)

// ValidationError represents a manifest validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
