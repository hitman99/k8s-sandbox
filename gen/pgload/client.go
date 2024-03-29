// Code generated by goa v3.0.2, DO NOT EDIT.
//
// pgload client
//
// Command:
// $ goa gen github.com/hitman99/k8s-sandbox/design

package pgload

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "pgload" service client.
type Client struct {
	LoadEndpoint goa.Endpoint
}

// NewClient initializes a "pgload" service client given the endpoints.
func NewClient(load goa.Endpoint) *Client {
	return &Client{
		LoadEndpoint: load,
	}
}

// Load calls the "load" endpoint of the "pgload" service.
func (c *Client) Load(ctx context.Context, p *LoadPayload) (res *JSONStatus, err error) {
	var ires interface{}
	ires, err = c.LoadEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*JSONStatus), nil
}
