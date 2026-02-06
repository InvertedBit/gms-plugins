package components

import (
	"testing"
)

func TestNewComponentViewModel_Create(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test-component", nil, nil, false, "")

	if vm.Name != "test-component" {
		t.Errorf("Expected Name 'test-component', got '%s'", vm.Name)
	}

	if vm.IsEdit != false {
		t.Error("Expected IsEdit to be false")
	}

	if vm.SubmitURL != "/admin/instances/test-component" {
		t.Errorf("Expected SubmitURL '/admin/instances/test-component', got '%s'", vm.SubmitURL)
	}

	if vm.CancelURL != "/admin/instances" {
		t.Errorf("Expected CancelURL '/admin/instances', got '%s'", vm.CancelURL)
	}

	if vm.FormErrors == nil {
		t.Error("FormErrors should not be nil")
	}

	if vm.Properties == nil {
		t.Error("Properties should not be nil")
	}

	if vm.Media == nil {
		t.Error("Media should not be nil")
	}
}

func TestNewComponentViewModel_Edit(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test-component", nil, nil, true, "")

	if vm.IsEdit != true {
		t.Error("Expected IsEdit to be true")
	}

	if vm.SubmitURL != "/admin/instances/test-component" {
		t.Errorf("Expected SubmitURL '/admin/instances/test-component', got '%s'", vm.SubmitURL)
	}
}

func TestComponentViewModel_GetProperty_Default(t *testing.T) {
	properties := []ComponentProperty{
		{Slug: "title", Key: "title", Value: "Hello", Type: Default},
		{Slug: "subtitle", Key: "subtitle", Value: "World", Type: Default},
	}

	vm := NewComponentViewModel("test-component-0", "test", properties, nil, false, "")

	result := vm.GetProperty("title")
	if result != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", result)
	}

	result = vm.GetProperty("subtitle")
	if result != "World" {
		t.Errorf("Expected 'World', got '%s'", result)
	}
}

func TestComponentViewModel_GetProperty_LayoutOverride(t *testing.T) {
	properties := []ComponentProperty{
		{Slug: "title", Key: "title", Value: "Default Title", Type: Default},
		{Slug: "title", Key: "title", Value: "Layout Title", Type: LayoutOverride},
	}

	vm := NewComponentViewModel("test-component-0", "test", properties, nil, false, "")

	result := vm.GetProperty("title")
	if result != "Layout Title" {
		t.Errorf("Expected 'Layout Title', got '%s'", result)
	}
}

func TestComponentViewModel_GetProperty_PageOverride(t *testing.T) {
	properties := []ComponentProperty{
		{Slug: "title", Key: "title", Value: "Default Title", Type: Default},
		{Slug: "title", Key: "title", Value: "Layout Title", Type: LayoutOverride},
		{Slug: "title", Key: "title", Value: "Page Title", Type: PageOverride},
	}

	vm := NewComponentViewModel("test-component-0", "test", properties, nil, false, "")

	result := vm.GetProperty("title")
	if result != "Page Title" {
		t.Errorf("Expected 'Page Title', got '%s'", result)
	}
}

func TestComponentViewModel_GetProperty_NotFound(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test", nil, nil, false, "")

	result := vm.GetProperty("nonexistent")
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestComponentViewModel_GetMediaURL(t *testing.T) {
	media := []ComponentMedia{
		{Slug: "avatar", FileName: "avatar.png", FileType: "image/png", URL: "/media/avatar.png"},
		{Slug: "banner", FileName: "banner.jpg", FileType: "image/jpeg", URL: "/media/banner.jpg"},
	}

	vm := NewComponentViewModel("test-component-0", "test", nil, media, false, "")

	result := vm.GetMediaURL("avatar")
	if result != "/media/avatar.png" {
		t.Errorf("Expected '/media/avatar.png', got '%s'", result)
	}

	result = vm.GetMediaURL("banner")
	if result != "/media/banner.jpg" {
		t.Errorf("Expected '/media/banner.jpg', got '%s'", result)
	}
}

func TestComponentViewModel_GetMediaURL_NotFound(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test", nil, nil, false, "")

	result := vm.GetMediaURL("nonexistent")
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestComponentViewModel_GetFormError(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test", nil, nil, false, "")
	vm.FormErrors["email"] = "Invalid email address"
	vm.FormErrors["password"] = "Password too short"

	result := vm.GetFormError("email")
	if result != "Invalid email address" {
		t.Errorf("Expected 'Invalid email address', got '%s'", result)
	}

	result = vm.GetFormError("password")
	if result != "Password too short" {
		t.Errorf("Expected 'Password too short', got '%s'", result)
	}
}

func TestComponentViewModel_GetFormError_NotFound(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test", nil, nil, false, "")

	result := vm.GetFormError("nonexistent")
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestComponentViewModel_GetRESTPath_ComponentPath(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test", nil, nil, false, "/custom/context")

	result := vm.GetRESTPath("/subquery", ComponentPath)
	if result != "/sys/components/test-component-0/subquery" {
		t.Errorf("Expected '/sys/components/test-component-0/subquery' for RESTPath in ComponentPath mode, got '%s'", result)
	}
}

func TestComponentViewModel_GetRESTPath_PagePath(t *testing.T) {
	vm := NewComponentViewModel("test-component-0", "test", nil, nil, false, "/custom/context")

	result := vm.GetRESTPath("/subquery", PagePath)
	if result != "/custom/context/subquery" {
		t.Errorf("Expected '/custom/context/subquery' for RESTPath in PagePath mode, got '%s'", result)
	}
}
