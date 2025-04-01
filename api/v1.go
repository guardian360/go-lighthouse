package api

import (
	"github.com/guardian360/go-lighthouse/client"
	"github.com/mitchellh/mapstructure"
)

// APIv1 is the representation of Lighthouse API v1. It is meant to be used as
// a namespace for all v1 API resources.
type APIv1 struct {
	Client *client.Client
}

// V1 creates a new APIv1 instance.
func V1(c *client.Client) *APIv1 {
	return &APIv1{Client: c}
}

// Heartbeat retrieves the heartbeat API.
func (api *APIv1) Heartbeat() *HeartbeatAPIv1 {
	return NewHeartbeatAPIv1(api.Client)
}

// Probes retrieves the probes API.
func (api *APIv1) Probes() *ProbesAPIv1 {
	return NewProbesAPIv1(api.Client)
}

// HackerAlertAppliances retrieves the hacker alert appliances API.
func (api *APIv1) HackerAlertAppliances() *HackerAlertAppliancesAPIv1 {
	return NewHackerAlertAppliancesAPIv1(api.Client)
}

// APIv1Response is the response wrapper for API v1.
type APIv1Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Decode decodes the Data field into the provided struct.
func (r *APIv1Response) Decode(v interface{}) error {
	return mapstructure.Decode(r.Data, v)
}
