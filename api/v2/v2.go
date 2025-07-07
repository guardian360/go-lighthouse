package v2

import (
	"github.com/guardian360/go-lighthouse/client"
	"github.com/mitchellh/mapstructure"
)

// API is the representation of Lighthouse API v2. It is meant to be used as
// a namespace for all v2 API resources.
type API struct {
	Client *client.Client
}

// New creates a new API instance.
func New(c *client.Client) *API {
	return &API{Client: c}
}

// Heartbeat retrieves the heartbeat API.
func (api *API) Heartbeat() *HeartbeatAPI {
	return NewHeartbeatAPI(api.Client)
}

// Probes retrieves the probes API.
func (api *API) Probes() *ProbesAPI {
	return NewProbesAPI(api.Client)
}

// ScanTrackers retrieves the scan trackers API.
func (api *API) ScanTrackers() *ScanTrackersAPI {
	return NewScanTrackersAPI(api.Client)
}

// APIResponse is the response wrapper for API v2.
type APIResponse struct {
	Data  interface{}      `json:"data"`
	Links APIResponseLinks `json:"links,omitempty"`
	Meta  APIResponseMeta  `json:"meta,omitempty"`
}

// APIResponseLinks contains pagination links for the API response.
type APIResponseLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

// APIResponseMeta contains metadata about the API response, such as pagination
// information.
type APIResponseMeta struct {
	CurrentPage int                   `json:"current_page"`
	From        int                   `json:"from"`
	LastPage    int                   `json:"last_page"`
	Links       []APIResponseMetaLink `json:"links"`
	Path        string                `json:"path"`
	PerPage     int                   `json:"per_page"`
	To          int                   `json:"to"`
	Total       int                   `json:"total"`
}

// APIResponseMetaLink represents a link in the metadata of the API response.
type APIResponseMetaLink struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

// Decode decodes the Data field into the provided struct.
func (r *APIResponse) Decode(v interface{}) error {
	return mapstructure.Decode(r, v)
}
