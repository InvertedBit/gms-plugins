package components

import (
	"os"
	"plugin"

	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

var ComponentRenderer *Renderer

type Renderer struct {
	Components map[string]Component
}

func NewRenderer() *Renderer {
	return &Renderer{
		Components: map[string]Component{
			"container": {
				Name:        "Container",
				Description: "A basic container (div) component",
				Render:      RenderContainerComponent,
				Children:    []Component{},
			},
			// Add more components here as needed
		},
	}
}

func (r *Renderer) TryLoadPlugins(pluginDir string) error {
	// Get file list in pluginDir
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
		getComponentsSymbol, err := plugin.Lookup("GetComponents")
		if err != nil {
			return err
		}
		getComponentsFunc := getComponentsSymbol.(func() map[string]Component)
		components := getComponentsFunc()
		for name, component := range components {
			if _, exists := r.Components[name]; exists {
				continue // Skip if component with same name already exists
			}
			r.Components[name] = component
		}
	}
	// For each file, attempt to load as plugin
	// If plugin has a Component, register it in r.Components
	return nil
}

func (r *Renderer) PrintLoadedComponents() {
	for name, component := range r.Components {
		println("Loaded component:", name, "-", component.Name)
	}
}

func (r *Renderer) RenderComponent(vm *ComponentViewModel) gomponents.Node {
	if component, exists := r.Components[vm.Name]; exists {
		return component.Render(vm)
	} else {
		return html.Div(
			gomponents.Text("Unknown Component"),
		)
	}
}
