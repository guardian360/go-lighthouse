package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HeartbeatAPIv2 is the v1 API for the heartbeat resource.
type HeartbeatAPI struct {
	api.APIRequestHandler
}

// NewHeartbeatAPI creates a new HeartbeatAPI instance.
func NewHeartbeatAPI(c *client.Client) *HeartbeatAPI {
	return &HeartbeatAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/heartbeat",
		},
	}
}

// Get retrieves the API heartbeat response.
func (h *HeartbeatAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}
