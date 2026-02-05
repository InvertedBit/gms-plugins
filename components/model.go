package components

import (
	"github.com/invertedbit/gms-plugins/hooks"
	"maragu.dev/gomponents"
)

type RenderFunc func(*ComponentViewModel) gomponents.Node

type ComponentPropertyType int

const (
	Default ComponentPropertyType = iota
	LayoutOverride
	PageOverride
)

type PropertyType string

const (
	PropertyString PropertyType = "string"
	PropertyInt    PropertyType = "int"
	PropertyBool   PropertyType = "bool"
	PropertyList   PropertyType = "list"
)

type ComponentProperty struct {
	Slug     string
	Key      string
	Value    string
	Type     ComponentPropertyType
	TypeStr  PropertyType
	Required bool
	Default  string
}

type ComponentMedia struct {
	Slug     string
	FileName string
	FileType string
	URL      string
}

type ComponentInterface interface {
	Init(properties []ComponentProperty, media []ComponentMedia) error
	Render(vm *ComponentViewModel) gomponents.Node
}

type Component struct {
	Slug        string
	Name        string
	Description string
	Properties  []ComponentProperty
	Media       []ComponentMedia
	Children    []Component
	Render      RenderFunc
	Hooks       map[string]hooks.Hook
}

func (c Component) Init(properties []ComponentProperty, media []ComponentMedia) error {
	// Default implementation does nothing, can be overridden by specific components
	c.Properties = properties
	c.Media = media
	c.Hooks = make(map[string]hooks.Hook)
	return nil
}
