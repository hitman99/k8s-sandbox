// Code generated by goa v3.0.2, DO NOT EDIT.
//
// pgload client HTTP transport
//
// Command:
// $ goa gen github.com/hitman99/k8s-sandbox/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the pgload service endpoint HTTP clients.
type Client struct {
	// Load Doer is the HTTP client used to make requests to the load endpoint.
	LoadDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the pgload service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		LoadDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Load returns an endpoint that makes HTTP requests to the pgload service load
// server.
func (c *Client) Load() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoadRequest(c.encoder)
		decodeResponse = DecodeLoadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoadDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("pgload", "load", err)
		}
		return decodeResponse(resp)
	}
}
