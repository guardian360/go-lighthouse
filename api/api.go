// Package api provides the API for the Lighthouse API client.
package api

import (
	"github.com/guardian360/go-lighthouse/client"
	"github.com/mitchellh/mapstructure"
)

// APIRequestHandler is the base resource for all API resources.
type APIRequestHandler struct {
	Client  *client.Client
	BaseURL string
}

func do[T any](r APIRequestHandler, method, url string, data map[string]interface{}) (*T, error) {
	resp, err := r.Client.Do(method, url, nil)
	if err != nil {
		return nil, err
	}

	var decodedResponse T
	if err := mapstructure.Decode(resp, &decodedResponse); err != nil {
		return nil, err
	}

	return &decodedResponse, nil
}
