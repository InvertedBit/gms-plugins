package main

import (
	"context"
	"fmt"
	"strconv"

	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"

	"github.com/invertedbit/gms-plugins/components"
	"github.com/invertedbit/gms-plugins/hooks"
	"github.com/invertedbit/gms-plugins/plugins"
)

// HelloWorldComponent is a simple example component
type HelloWorldComponent struct {
	components.Component
	Message string
}

func (hwc HelloWorldComponent) Init(properties []components.ComponentProperty, media []components.ComponentMedia) error {
	hwc.Component.Init(properties, media)
	// Initialize component properties if needed
	for _, prop := range properties {
		if prop.Key == "message" {
			if prop.TypeStr == components.PropertyString {
				hwc.Message = prop.Value
			} else {
				return fmt.Errorf("invalid type for message property, expected string, got %s", prop.TypeStr)
			}
		}
	}
	return nil
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
	components.Component
	InitialCount int
}

func (cc CounterComponent) Init(properties []components.ComponentProperty, media []components.ComponentMedia) error {
	cc.Component.Init(properties, media)
	// Initialize component properties if needed
	for _, prop := range properties {
		if prop.Key == "initialCount" {
			if prop.TypeStr == components.PropertyInt {
				// Convert string to int
				var err error
				cc.InitialCount, err = strconv.Atoi(prop.Value)
				if err != nil {
					return fmt.Errorf("invalid value for initialCount property, expected integer, got %s", prop.Value)
				}
			} else {
				return fmt.Errorf("invalid type for initialCount property, expected int, got %s", prop.TypeStr)
			}
		}
	}
	cc.Hooks["counterUpdated"] = hooks.Hook{
		Name:        "counterUpdated",
		Description: "Handles counter updates from the client",
		Priority:    0,
		Handler:     cc.HandleCounterUpdate,
	}
	return nil
}

func (cc CounterComponent) HandleCounterUpdate(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	// Handle counter update logic here, e.g. increment or decrement based on args
	action, ok := args["action"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid action argument")
	}

	currentCount := cc.InitialCount
	if action == "increment" {
		currentCount++
	} else if action == "decrement" {
		currentCount--
	} else {
		return nil, fmt.Errorf("invalid action: %s", action)
	}

	// Return the updated count to be sent back to the client
	return map[string]interface{}{
		"count": currentCount,
	}, nil
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
				htmx.Put(vm.RestContext+"/counter/decrement"),
				gomponents.Text("-"),
			),
			html.Span(
				html.Class("text-2xl font-mono w-12 text-center"),
				gomponents.Text(strconv.Itoa(cc.InitialCount)),
			),
			html.Button(
				html.Class("btn btn-primary"),
				htmx.Put(vm.RestContext+"/counter/increment"),
				htmx.Target("closest .counter"),
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
			Components: map[string]components.ComponentInterface{
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
