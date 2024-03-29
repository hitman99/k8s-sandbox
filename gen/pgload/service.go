// Code generated by goa v3.0.2, DO NOT EDIT.
//
// pgload service
//
// Command:
// $ goa gen github.com/hitman99/k8s-sandbox/design

package pgload

import (
	"context"

	pgloadviews "github.com/hitman99/k8s-sandbox/gen/pgload/views"
)

// The pgload service performs operations on postgres database.
type Service interface {
	// Load implements load.
	Load(context.Context, *LoadPayload) (res *JSONStatus, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "pgload"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"load"}

// LoadPayload is the payload type of the pgload service load method.
type LoadPayload struct {
	// How many records to generate in the table
	Count int
}

// JSONStatus is the result type of the pgload service load method.
type JSONStatus struct {
	// result code
	Code uint
	// status info
	Status string
	// processing time
	Time *string
}

// NewJSONStatus initializes result type JSONStatus from viewed result type
// JSONStatus.
func NewJSONStatus(vres *pgloadviews.JSONStatus) *JSONStatus {
	var res *JSONStatus
	switch vres.View {
	case "default", "":
		res = newJSONStatus(vres.Projected)
	}
	return res
}

// NewViewedJSONStatus initializes viewed result type JSONStatus from result
// type JSONStatus using the given view.
func NewViewedJSONStatus(res *JSONStatus, view string) *pgloadviews.JSONStatus {
	var vres *pgloadviews.JSONStatus
	switch view {
	case "default", "":
		p := newJSONStatusView(res)
		vres = &pgloadviews.JSONStatus{p, "default"}
	}
	return vres
}

// newJSONStatus converts projected type JSONStatus to service type JSONStatus.
func newJSONStatus(vres *pgloadviews.JSONStatusView) *JSONStatus {
	res := &JSONStatus{
		Time: vres.Time,
	}
	if vres.Code != nil {
		res.Code = *vres.Code
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	return res
}

// newJSONStatusView projects result type JSONStatus to projected type
// JSONStatusView using the "default" view.
func newJSONStatusView(res *JSONStatus) *pgloadviews.JSONStatusView {
	vres := &pgloadviews.JSONStatusView{
		Code:   &res.Code,
		Status: &res.Status,
		Time:   res.Time,
	}
	return vres
}
