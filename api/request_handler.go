package api

import (
	"net/url"

	"github.com/guardian360/go-lighthouse/client"
	"github.com/mitchellh/mapstructure"
)

// APIRequestHandler is the base resource for all API resources.
type APIRequestHandler struct {
	Client  *client.Client
	BaseURL string
	params  url.Values
}

func (r *APIRequestHandler) setParam(param, value string) {
	if r.params == nil {
		r.params = url.Values{}
	}
	r.params.Set(param, value)
}

func (r *APIRequestHandler) buildURL() string {
	if r.params == nil {
		return r.BaseURL
	}
	return r.BaseURL + "?" + r.params.Encode()
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
