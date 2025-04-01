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
			Client:  api.Client,
			Version: "v2",
			Path:    "probes",
		},
	}
}

// ScanTrackers retrieves the scan trackers API.
func (api *APIv2) ScanTrackers() *ScanTrackersAPI {
	return &ScanTrackersAPI{
		APIResource: APIResource{
			Client:  api.Client,
			Version: "v2",
			Path:    "scan-trackers",
		},
	}
}
