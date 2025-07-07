package v1

import (
	"github.com/guardian360/go-lighthouse/client"
)

// API is the representation of Lighthouse API v1. It is meant to be used as
// a namespace for all v1 API resources.
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

// HackerAlertAppliances retrieves the hacker alert appliances API.
func (api *API) HackerAlertAppliances() *HackerAlertAppliancesAPI {
	return NewHackerAlertAppliancesAPI(api.Client)
}

// ScanObjects retrieves the scan objects API.
func (api *API) ScanObjects() *ScanObjectsAPI {
	return NewScanObjectsAPI(api.Client)
}

// APIResponse is the response wrapper for API v1.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
