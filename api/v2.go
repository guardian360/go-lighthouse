package api

import (
	"github.com/guardian360/go-lighthouse/client"
	"github.com/mitchellh/mapstructure"
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

// Heartbeat retrieves the heartbeat API.
func (api *APIv2) Heartbeat() *HeartbeatAPIv2 {
	return NewHeartbeatAPIv2(api.Client)
}

// Probes retrieves the probes API.
func (api *APIv2) Probes() *ProbesAPIv2 {
	return NewProbesAPIv2(api.Client)
}

// ScanTrackers retrieves the scan trackers API.
func (api *APIv2) ScanTrackers() *ScanTrackersAPIv2 {
	return NewScanTrackersAPIv2(api.Client)
}

// APIv2Response is the response wrapper for API v2.
type APIv2Response struct {
	Data  interface{} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links,omitempty"`
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
	} `json:"meta,omitempty"`
}

// Decode decodes the Data field into the provided struct.
func (r *APIv2Response) Decode(v interface{}) error {
	return mapstructure.Decode(r.Data, v)
}
