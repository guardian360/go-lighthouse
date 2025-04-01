package api

import (
	"github.com/guardian360/go-lighthouse/client"
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
func (api *APIv1) Heartbeat() *HeartbeatAPI {
	return &HeartbeatAPI{
		APIResource: APIResource{
			Client: api.Client,
			Path:   "heartbeat",
		},
	}
}

// Probes retrieves the probes API.
func (api *APIv1) Probes() *ProbesAPI {
	return &ProbesAPI{
		APIResource: APIResource{
			Client:  api.Client,
			Version: "v1",
			Path:    "probes",
		},
	}
}

// HackerAlertAppliances retrieves the hacker alert appliances API.
func (api *APIv1) HackerAlertAppliances() *HackerAlertAppliancesAPI {
	return &HackerAlertAppliancesAPI{
		APIResource: APIResource{
			Client:  api.Client,
			Version: "v1",
			Path:    "hacker-alert-appliances",
		},
	}
}
