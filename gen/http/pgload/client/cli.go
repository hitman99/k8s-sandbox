// Code generated by goa v3.0.2, DO NOT EDIT.
//
// pgload HTTP client CLI support package
//
// Command:
// $ goa gen github.com/hitman99/k8s-sandbox/design

package client

import (
	"encoding/json"
	"fmt"

	pgload "github.com/hitman99/k8s-sandbox/gen/pgload"
)

// BuildLoadPayload builds the payload for the pgload load endpoint from CLI
// flags.
func BuildLoadPayload(pgloadLoadBody string) (*pgload.LoadPayload, error) {
	var err error
	var body LoadRequestBody
	{
		err = json.Unmarshal([]byte(pgloadLoadBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"count\": 8606277519841735856\n   }'")
		}
	}
	v := &pgload.LoadPayload{
		Count: body.Count,
	}
	return v, nil
}
