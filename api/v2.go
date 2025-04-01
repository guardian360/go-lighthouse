package api

import (
	"github.com/guardian360/go-lighthouse/client"
)

// APIv2 is the representation of Lighthouse API v2. It is meant to be used as
// a namespace for all v2 API resources.
type APIv2 struct {
	Client *client.Client
}

// V2 creates a new APIv2 instance.
func V2(c *client.Client) *APIv2 {
	return &APIv2{Client: c}
}

// Probes retrieves the probes API.
func (api *APIv2) Probes() *ProbesAPI {
	return &ProbesAPI{
		APIResource: APIResource{
			Client:      api.Client,
			Version:     "v2",
			Path:        "probes",
			APIResponse: &APIv2Response{},
		},
	}
}

// ScanTrackers retrieves the scan trackers API.
func (api *APIv2) ScanTrackers() *ScanTrackersAPI {
	return &ScanTrackersAPI{
		APIResource: APIResource{
			Client:      api.Client,
			Version:     "v2",
			Path:        "scan-trackers",
			APIResponse: &APIv2Response{},
		},
	}
}

// APIv2Response is the response wrapper for API v2.
type APIv2Response struct {
	Data  interface{} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int `json:"current_page"`
		From        int `json:"from"`
		LastPage    int `json:"last_page"`
		Links       []struct {
			URL    string `json:"url"`
			Label  string `json:"label"`
			Active bool   `json:"active"`
		} `json:"links"`
		Path    string `json:"path"`
		PerPage int    `json:"per_page"`
		To      int    `json:"to"`
		Total   int    `json:"total"`
	} `json:"meta"`
}

// Wrap wraps the response from the API.
func (r *APIv2Response) Wrap(resp map[string]interface{}) error {
	return wrap(resp, r)
}
