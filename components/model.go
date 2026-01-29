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

type ComponentProperty struct {
	Slug  string
	Key   string
	Value string
	Type  ComponentPropertyType
}

type ComponentMedia struct {
	Slug     string
	FileName string
	FileType string
	URL      string
}

type Component struct {
	Name        string
	Description string
	Children    []Component
	Render      RenderFunc
}
