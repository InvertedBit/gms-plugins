package components

type ComponentViewModel struct {
	IsEdit      bool
	SubmitURL   string
	CancelURL   string
	FormErrors  map[string]string
	Name        string
	Properties  map[string]ComponentProperty
	Media       map[string]ComponentMedia
	RestContext string
}

func NewComponentViewModel(name string, properties []ComponentProperty, media []ComponentMedia, isEdit bool, restContext string) *ComponentViewModel {
	submitURL := "/admin/instances/"
	if isEdit {
		submitURL = "/admin/instances/" + name
	}
	cancelURL := "/admin/instances"

	return &ComponentViewModel{
		IsEdit:      isEdit,
		SubmitURL:   submitURL,
		CancelURL:   cancelURL,
		FormErrors:  make(map[string]string),
		Name:        name,
		Properties:  make(map[string]ComponentProperty),
		Media:       make(map[string]ComponentMedia),
		RestContext: restContext,
	}
}

func (vm *ComponentViewModel) GetProperty(slug string) string {
	defaultProperty := ""
	layoutOverride := ""
	for _, prop := range vm.Properties {
		if prop.Slug == slug {
			switch prop.Type {
			case PageOverride:
				return prop.Value
			case LayoutOverride:
				layoutOverride = prop.Value
			case Default:
				defaultProperty = prop.Value
			}
		}
	}
	if layoutOverride != "" {
		return layoutOverride
	}
	return defaultProperty
}

func (vm *ComponentViewModel) GetMediaURL(slug string) string {
	for _, media := range vm.Media {
		if media.Slug == slug {
			return media.URL
		}
	}
	return ""
}

func (vm *ComponentViewModel) GetFormError(field string) string {
	if err, exists := vm.FormErrors[field]; exists {
		return err
	}
	return ""
}
