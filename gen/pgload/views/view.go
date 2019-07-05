// Code generated by goa v3.0.2, DO NOT EDIT.
//
// pgload views
//
// Command:
// $ goa gen github.com/hitman99/k8s-sandbox/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// JSONStatus is the viewed result type that is projected based on a view.
type JSONStatus struct {
	// Type to project
	Projected *JSONStatusView
	// View to render
	View string
}

// JSONStatusView is a type that runs validations on a projected type.
type JSONStatusView struct {
	// result code
	Code *uint
	// status info
	Status *string
	// processing time
	Time *string
}

var (
	// JSONStatusMap is a map of attribute names in result type JSONStatus indexed
	// by view name.
	JSONStatusMap = map[string][]string{
		"default": []string{
			"code",
			"status",
			"time",
		},
	}
)

// ValidateJSONStatus runs the validations defined on the viewed result type
// JSONStatus.
func ValidateJSONStatus(result *JSONStatus) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJSONStatusView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateJSONStatusView runs the validations defined on JSONStatusView using
// the "default" view.
func ValidateJSONStatusView(result *JSONStatusView) (err error) {
	if result.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	return
}
