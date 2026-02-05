package main

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"

	"github.com/invertedbit/gms-plugins/components"
	"github.com/invertedbit/gms-plugins/plugins"
)

// HelloWorldComponent is a simple example component
type HelloWorldComponent struct {
	Message string
}

func (hwc HelloWorldComponent) Render(vm *components.ComponentViewModel) gomponents.Node {
	return html.Div(
		html.Class("hello-world p-4 bg-base-200 rounded-lg"),
		html.H2(
			html.Class("text-xl font-bold mb-2"),
			gomponents.Text("Hello World!"),
		),
		html.P(
			html.Class("text-base-content"),
			gomponents.Text(hwc.Message),
		),
	)
}

// CounterComponent demonstrates interactive state
type CounterComponent struct {
	InitialCount int
}

func (cc CounterComponent) Render(vm *components.ComponentViewModel) gomponents.Node {
	return html.Div(
		html.Class("counter p-4 border rounded-lg"),
		html.H3(
			html.Class("font-bold mb-2"),
			gomponents.Text("Counter Component"),
		),
		html.Div(
			html.Class("flex items-center gap-2"),
			html.Button(
				html.Class("btn btn-primary"),
				html.Data("onClick", "decrement()"),
				gomponents.Text("-"),
			),
			html.Span(
				html.Class("text-2xl font-mono w-12 text-center"),
				html.Data("hx-get", "/counter/value"),
				html.Data("hx-trigger", "counterUpdated from:body"),
				gomponents.Text("0"),
			),
			html.Button(
				html.Class("btn btn-primary"),
				html.Data("onClick", "increment()"),
				gomponents.Text("+"),
			),
		),
	)
}

// GetPlugins is the entry point for GMS plugin loading
func GetPlugins() map[string]plugins.Plugin {
	return map[string]plugins.Plugin{
		"hello-world": {
			Name:        "Hello World",
			Author:      "GMS Team",
			Version:     "1.0.0",
			Description: "A simple hello world component example",
			Components: map[string]components.Component{
				"hello-world": HelloWorldComponent{
					Message: "Welcome to the GMS Plugin System!",
				},
				"counter": CounterComponent{
					InitialCount: 0,
				},
			},
		},
	}
}
