# GMS Plugins

This repository contains the plugin system for the GMS (Go Modular System) CMS.

## Overview

GMS uses a plugin architecture that allows you to extend the system with custom components. Plugins are compiled as shared objects (.so files) and loaded at runtime.

## Quick Start

### 1. Create a new plugin package

```bash
mkdir my-plugin
cd my-plugin
go mod init github.com/yourusername/my-plugin
```

### 2. Import the plugin interfaces

```go
package main

import (
    "github.com/invertedbit/gms-plugins/components"
    "github.com/invertedbit/gms-plugins/plugins"
)
```

### 3. Define your components

Components are defined using the `Component` struct:

```go
import "maragu.dev/gomponents"
import "maragu.dev/gomponents/html"

type MyComponent struct {
    Title string
}

func (mc MyComponent) Render(vm *components.ComponentViewModel) gomponents.Node {
    return html.Div(
        html.Class("my-component"),
        html.Text(mc.Title),
    )
}
```

### 4. Export your plugins

Create a `GetPlugins()` function that returns a map of plugins:

```go
func GetPlugins() map[string]plugins.Plugin {
    return map[string]plugins.Plugin{
        "my-plugin": {
            Name:        "My Plugin",
            Author:      "Your Name",
            Version:     "1.0.0",
            Description: "A sample plugin for GMS",
            Components: map[string]components.Component{
                "my-component": MyComponent{},
            },
        },
    }
}
```

### 5. Build as a shared library

Add this to your `go.mod`:

```go
// Build as plugin
go build -buildmode=plugin -o my-plugin.so .
```

### 6. Place in plugins directory

Copy the `.so` file to your GMS plugins directory:

```bash
cp my-plugin.so /path/to/gms/plugins/
```

## Component Properties

Components can have configurable properties:

```go
type ConfigurableComponent struct {
    Title    string
    Subtitle string
}

func (cc ConfigurableComponent) Render(vm *components.ComponentViewModel) gomponents.Node {
    // Access properties from view model
    title := vm.GetProperty("title")
    subtitle := vm.GetProperty("subtitle")

    return html.Div(
        html.H1(html.Text(title)),
        html.P(html.Text(subtitle)),
    )
}
```

Available property types:
- `Default` - Standard property value
- `LayoutOverride` - Override for layout context
- `PageOverride` - Override for page context

## Example Components

This package includes example components in `components/examples.go`:

| Component | Description |
|-----------|-------------|
| `ButtonComponent` | Configurable button with variants |
| `CardComponent` | Card container with image support |
| `AlertComponent` | Alert notifications with types |
| `InputComponent` | Form input with validation |
| `BadgeComponent` | Status badges with colors |

## Plugin Manager

The `PluginManager` handles loading and managing plugins:

```go
pm := &plugins.PluginManager{}

// Load all plugins from a directory
err := pm.TryLoadPlugins("./plugins/")

// Get loaded plugins
plugins := pm.GetLoadedPlugins()
```

## API Reference

### Component

```go
type Component struct {
    Name        string              // Display name
    Description string              // Component description
    Children    []Component         // Child components
    Render      RenderFunc          // Render function
}
```

### ComponentProperty

```go
type ComponentProperty struct {
    Slug  string               // Property identifier
    Key   string               // Property key
    Value string               // Property value
    Type  ComponentPropertyType // Property type (Default, LayoutOverride, PageOverride)
}
```

### ComponentMedia

```go
type ComponentMedia struct {
    Slug     string // Media identifier
    FileName string // Original filename
    FileType string // MIME type
    URL      string // Access URL
}
```

### ComponentViewModel

```go
type ComponentViewModel struct {
    IsEdit     bool                     // Edit mode flag
    SubmitURL  string                   // Form submission URL
    CancelURL  string                   // Cancel URL
    FormErrors map[string]string        // Form validation errors
    Name       string                   // Component name
    Properties map[string]ComponentProperty // Component properties
    Media      map[string]ComponentMedia    // Component media
}
```

## Best Practices

1. **Use semantic versioning** for your plugins
2. **Document your components** with clear descriptions
3. **Handle errors gracefully** in your Render functions
4. **Test your components** before distribution
5. **Follow GMS styling conventions** (Tailwind CSS)

## Directory Structure

```
gms-plugins/
├── components/          # Component interfaces and types
│   ├── examples.go     # Example components
│   ├── model.go        # Core component models
│   ├── renderer.go     # Renderer interface
│   └── viewmodel.go    # View model implementation
├── plugins/             # Plugin loading and management
│   ├── general.go      # Plugin manager implementation
│   └── plugin_manager_test.go
├── main.go             # Package entry point
├── go.mod              # Go module definition
└── README.md           # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add your components or improvements
4. Submit a pull request

## License

MIT License
