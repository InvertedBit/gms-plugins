package components

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// ButtonComponent is a simple button component example
type ButtonComponent struct {
	Label    string
	Variant  string // "primary", "secondary", "outline"
	Disabled bool
	OnClick  string
}

func (bc ButtonComponent) Render(vm *ComponentViewModel) gomponents.Node {
	class := "btn"
	if bc.Variant != "" {
		class += " btn-" + bc.Variant
	}
	if bc.Disabled {
		class += " disabled"
	}

	return html.Button(
		html.Class(class),
		html.Type("button"),
		gomponents.Text(bc.Label),
	)
}

// CardComponent is a card container component example
type CardComponent struct {
	Title     string
	Content   string
	ImageURL  string
	ImageAlt  string
	Footer    string
	Clickable bool
}

func (cc CardComponent) Render(vm *ComponentViewModel) gomponents.Node {
	cardClass := "card"
	if cc.Clickable {
		cardClass += " card-bordered cursor-pointer hover:shadow-lg transition-shadow"
	}

	children := []gomponents.Node{}

	if cc.ImageURL != "" {
		children = append(children,
			html.Figure(
				html.Class("relative"),
				html.Img(
					html.Src(cc.ImageURL),
					html.Alt(cc.ImageAlt),
					html.Class("w-full h-48 object-cover"),
				),
			),
		)
	}

	children = append(children,
		html.Div(
			html.Class("card-body"),
			html.H2(
				html.Class("card-title"),
				gomponents.Text(cc.Title),
			),
			html.P(
				gomponents.Text(cc.Content),
			),
		),
	)

	if cc.Footer != "" {
		children = append(children,
			html.Div(
				html.Class("card-footer justify-end"),
				gomponents.Text(cc.Footer),
			),
		)
	}

	return html.Div(
		html.Class(cardClass),
		gomponents.Group(children),
	)
}

// AlertComponent is an alert/notification component example
type AlertComponent struct {
	Message     string
	Type        string // "info", "success", "warning", "error"
	Dismissible bool
	Title       string
}

func (ac AlertComponent) Render(vm *ComponentViewModel) gomponents.Node {
	alertClass := "alert"
	if ac.Type != "" {
		alertClass += " alert-" + ac.Type
	}

	children := []gomponents.Node{}

	if ac.Title != "" {
		children = append(children,
			html.H3(
				html.Class("font-bold"),
				gomponents.Text(ac.Title),
			),
		)
	}

	children = append(children,
		html.Div(
			gomponents.Text(ac.Message),
		),
	)

	if ac.Dismissible {
		children = append(children,
			html.Button(
				html.Class("btn btn-sm btn-ghost"),
				gomponents.Text("×"),
			),
		)
	}

	return html.Div(
		html.Class(alertClass),
		gomponents.Group(children),
	)
}

// InputComponent is a form input component example
type InputComponent struct {
	Name        string
	Label       string
	Placeholder string
	Type        string // "text", "email", "password", "number"
	Required    bool
	Value       string
	HelpText    string
	Error       string
}

func (ic InputComponent) Render(vm *ComponentViewModel) gomponents.Node {
	inputClass := "input input-bordered w-full"
	if ic.Error != "" {
		inputClass += " input-error"
	}

	children := []gomponents.Node{
		html.Label(
			html.Class("label"),
			html.Span(
				html.Class("label-text"),
				gomponents.Text(ic.Label),
			),
		),
		html.Input(
			html.Type(ic.Type),
			html.Name(ic.Name),
			html.Placeholder(ic.Placeholder),
			html.Class(inputClass),
			html.Value(ic.Value),
		),
	}

	if ic.HelpText != "" {
		children = append(children,
			html.Label(
				html.Class("label"),
				html.Span(
					html.Class("label-text-alt"),
					gomponents.Text(ic.HelpText),
				),
			),
		)
	}

	if ic.Error != "" {
		children = append(children,
			html.Label(
				html.Class("label"),
				html.Span(
					html.Class("label-text-alt text-error"),
					gomponents.Text(ic.Error),
				),
			),
		)
	}

	return html.Div(
		html.Class("form-control w-full"),
		gomponents.Group(children),
	)
}

// BadgeComponent is a status badge component example
type BadgeComponent struct {
	Label   string
	Color   string // "neutral", "primary", "secondary", "accent", "ghost"
	Size    string // "xs", "sm", "md", "lg"
	Rounded bool
}

func (bc BadgeComponent) Render(vm *ComponentViewModel) gomponents.Node {
	badgeClass := "badge"
	if bc.Color != "" {
		badgeClass += " badge-" + bc.Color
	}
	if bc.Size != "" {
		badgeClass += " badge-" + bc.Size
	}
	if bc.Rounded {
		badgeClass += " badge-rounded"
	}

	return html.Span(
		html.Class(badgeClass),
		gomponents.Text(bc.Label),
	)
}
