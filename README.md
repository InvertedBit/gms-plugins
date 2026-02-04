# GMS Plugins

This repository contains the plugin system for the GMS (Go Modular System) CMS.

## Overview

GMS uses a plugin architecture that allows you to extend the system with custom components and hooks. Plugins are compiled as shared objects (.so files) and distributed as ZIP archives with a manifest.json file.

## Quick Start

### 1. Create a new plugin package

```bash
mkdir my-plugin
cd my-plugin
go mod init github.com/yourusername/my-plugin
```

### 2. Create manifest.json

Every plugin must have a `manifest.json` file:

```json
{
  "name": "my-plugin",
  "version": "1.0.0",
  "author": "Your Name",
  "description": "A sample plugin for GMS",
  "gmsVersion": ">=1.0.0",
  "permissions": ["components"],
  "hooks": [],
  "components": [
    {
      "slug": "my-component",
      "name": "My Component",
      "description": "A custom component"
    }
  ]
}
```

### 3. Define your components

Components are defined using the `Component` struct:

```go
import "maragu.dev/gomponents"
import "maragu.dev/gomponents/html"

type MyComponent struct{}

func (mc MyComponent) Render(vm *components.ComponentViewModel) gomponents.Node {
    return html.Div(
        html.Class("my-component p-4 bg-base-200 rounded"),
        html.Text("Hello from my plugin!"),
    )
}
```

### 4. Define hooks (optional)

Plugins can register hooks for GMS lifecycle events:

```go
func OnPageRender(ctx context.Context, args map[string]interface{}) error {
    // Modify page before rendering
    return nil
}
```

### 5. Export your plugins

Create a `GetPlugins()` function that returns a map of plugins:

```go
import (
    "context"

    "github.com/invertedbit/gms-plugins/components"
    "github.com/invertedbit/gms-plugins/plugins"
)

func GetPlugins() map[string]plugins.Plugin {
    return map[string]plugins.Plugin{
        "my-plugin": {
            Name:        "My Plugin",
            Author:      "Your Name",
            Version:     "1.0.0",
            Description: "A sample plugin for GMS",
            Permissions: []string{"components"},
            Components: map[string]components.Component{
                "my-component": MyComponent{},
            },
            Hooks: map[string]plugins.HookHandler{
                "onPageRender": OnPageRender,
            },
        },
    }
}
```

### 6. Build and package

Use the Makefile to build and package your plugin:

```bash
# Build plugin (runs tests first, generates checksum)
make build

# Output: my-plugin-1.0.0.zip
```

## Plugin Packaging

### ZIP Structure

```
my-plugin-1.0.0.zip
├── manifest.json          # Plugin metadata
├── plugin.so              # Compiled Go plugin
└── assets/                # Optional static assets
    ├── css/styles.css
    └── js/scripts.js
```

### Build Commands

```bash
# Run tests only
make test

# Build plugin with checksum
make build

# Build + checksum + GPG signature
make sign

# Run tests, build, checksum, and sign
make all

# Clean all build artifacts
make clean
```

### Verification

Users can verify plugin integrity:

```bash
# Verify checksum
sha256sum -c my-plugin-1.0.0.zip.sha256

# Verify GPG signature (if signed)
gpg --verify my-plugin-1.0.0.zip.asc my-plugin-1.0.0.zip
```

## Hook System

### Available Hooks

| Hook Name | Description | Arguments |
|-----------|-------------|-----------|
| `onStartup` | Called when GMS starts | - |
| `onShutdown` | Called before GMS shuts down | - |
| `onPageRender` | Called before page rendering | `page: Page, context: map` |
| `onPageRendered` | Called after page rendering | `page: Page, html: string` |
| `onUserLogin` | Called after successful login | `user: User, session: Session` |
| `onUserLogout` | Called on logout | `user: User` |
| `onMediaUpload` | Called after media upload | `media: Media` |
| `onComponentRender` | Called before component render | `component: Component, vm: ComponentViewModel` |

### Hook Handler Signature

```go
type HookHandler func(ctx context.Context, args map[string]interface{}) error
```

### Declaring Hooks in manifest.json

```json
{
  "hooks": [
    {
      "name": "onPageRender",
      "priority": 100
    }
  ]
}
```

## Component Properties

Components can have configurable properties defined in manifest.json:

```json
{
  "components": [
    {
      "slug": "my-component",
      "name": "My Component",
      "description": "A custom component",
      "properties": [
        {
          "slug": "title",
          "type": "string",
          "required": false,
          "default": "Hello"
        }
      ]
    }
  ]
}
```

Access properties in your component:

```go
func (mc MyComponent) Render(vm *components.ComponentViewModel) gomponents.Node {
    title := vm.GetProperty("title")
    return html.Div(
        html.Text(title),
    )
}
```

## Example Components

This package includes example components in `components/examples.go`:

| Component | Description |
|-----------|-------------|
| `ButtonComponent` | Configurable button with variants |
| `CardComponent` | Card container with image support |
| `AlertComponent` | Alert notifications with types |
| `InputComponent` | Form input with validation |
| `BadgeComponent` | Status badges with colors |

## API Reference

### Plugin

```go
type Plugin struct {
    Name        string                       // Display name
    Author      string                       // Plugin author
    Version     string                       // Semantic version
    Description string                       // Plugin description
    Permissions []string                     // Required permissions
    Components  map[string]Component         // Component definitions
    Hooks       map[string]HookHandler      // Hook handlers
}
```

### Component

```go
type Component struct {
    Slug        string              // Unique identifier
    Name        string              // Display name
    Description string              // Component description
    Properties  []ComponentProperty  // Configurable properties
    Children    []Component         // Child components
    Render      RenderFunc          // Render function
}
```

### HookHandler

```go
type HookHandler func(ctx context.Context, args map[string]interface{}) error
```

### ComponentViewModel

```go
type ComponentViewModel struct {
    IsEdit     bool                          // Edit mode flag
    SubmitURL  string                        // Form submission URL
    CancelURL  string                        // Cancel URL
    FormErrors map[string]string             // Form validation errors
    Name       string                        // Component name
    Properties map[string]ComponentProperty   // Component properties
    Media      map[string]ComponentMedia      // Component media
}
```

## Directory Structure

```
gms-plugins/
├── .github/
│   └── workflows/
│       ├── ci.yml         # CI pipeline
│       └── release.yml    # Release pipeline
├── components/
│   ├── examples.go       # Example components
│   ├── model.go          # Core component models
│   ├── renderer.go       # Renderer interface
│   └── viewmodel.go      # View model implementation
├── manifest/
│   └── manifest.go       # Manifest parsing and validation
├── plugins/
│   └── general.go        # Plugin manager and hook registry
├── Makefile              # Build and packaging
├── manifest.json         # Example manifest
└── README.md            # This file
```

## CI/CD

### GitHub Actions

- **CI Workflow**: Runs tests on every push/PR
- **Release Workflow**: Builds and releases on tag creation

### Automated Checks

1. All tests must pass before build
2. SHA-256 checksum generated for every release
3. Optional GPG signature for authenticated releases

## Best Practices

1. **Use semantic versioning** for your plugins
2. **Document your components** with clear descriptions
3. **Define hooks** for extensibility points
4. **Handle errors gracefully** in Render and Hook functions
5. **Test your components** before distribution
6. **Follow GMS styling conventions** (Tailwind CSS)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add your components, hooks, or improvements
4. Ensure all tests pass (`make test`)
5. Submit a pull request

## License

MIT License
