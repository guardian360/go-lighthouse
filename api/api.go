// Package api provides the API for the Lighthouse API client.
package api

import (
	"net/url"

	"github.com/guardian360/go-lighthouse/client"
	"github.com/mitchellh/mapstructure"
)

// APIRequestPayload is the payload for API requests. It contains the data to
// be sent in the request body.
type APIRequestPayload map[string]interface{}

// APIRequestHandler is the base resource for all API resources.
type APIRequestHandler struct {
	Client  *client.Client
	BaseURL string
	params  url.Values
}

// SetParam sets a query parameter for the API request.
func (r *APIRequestHandler) SetParam(param, value string) {
	if r.params == nil {
		r.params = url.Values{}
	}
	r.params.Set(param, value)
}

// BuildURL constructs the full URL for the API request, including any query
// parameters.
func (r *APIRequestHandler) BuildURL() string {
	if r.params == nil {
		return r.BaseURL
	}
	return r.BaseURL + "?" + r.params.Encode()
}

// Do executes an API request with the specified method, URL, and payload.
func Do[T any](r APIRequestHandler, method, url string, data APIRequestPayload) (*T, error) {
	resp, err := r.Client.Do(method, url, data)
	if err != nil {
		return nil, err
	}
	var decoded T
	if err := mapstructure.Decode(resp, &decoded); err != nil {
		return nil, err
	}
	return &decoded, nil
}
