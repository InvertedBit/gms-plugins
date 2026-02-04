package components

import (
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

type Component struct {
	Slug        string
	Name        string
	Description string
	Properties  []ComponentProperty
	Children    []Component
	Render      RenderFunc
}
